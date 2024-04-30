package record

import (
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/aler9/writerseeker"
	"github.com/bluenviron/mediacommon/pkg/formats/fmp4"

	"github.com/noranetworks/kvds/internal/logger"
)

func writePart(
	f io.Writer,
	sequenceNumber uint32,
	partTracks map[*formatFMP4Track]*fmp4.PartTrack,
) error {
	fmp4PartTracks := make([]*fmp4.PartTrack, len(partTracks))
	i := 0
	for _, partTrack := range partTracks {
		fmp4PartTracks[i] = partTrack
		i++
	}

	part := &fmp4.Part{
		SequenceNumber: sequenceNumber,
		Tracks:         fmp4PartTracks,
	}

	var ws writerseeker.WriterSeeker
	err := part.Marshal(&ws)
	if err != nil {
		return err
	}

	_, err = f.Write(ws.Bytes())
	return err
}

type formatFMP4Part struct {
	s              *formatFMP4Segment
	sequenceNumber uint32
	startDTS       time.Duration

	created    time.Time
	partTracks map[*formatFMP4Track]*fmp4.PartTrack
	endDTS     time.Duration
}

func newFormatFMP4Part(
	s *formatFMP4Segment,
	sequenceNumber uint32,
	startDTS time.Duration,
) *formatFMP4Part {
	return &formatFMP4Part{
		s:              s,
		startDTS:       startDTS,
		sequenceNumber: sequenceNumber,
		created:        timeNow(),
		partTracks:     make(map[*formatFMP4Track]*fmp4.PartTrack),
	}
}

func (p *formatFMP4Part) close() error {
	if p.s.fi == nil {
		p.s.path = path(p.created).encode(p.s.f.a.pathFormat)
		p.s.f.a.agent.Log(logger.Debug, "creating segment %s", p.s.path)

		err := os.MkdirAll(filepath.Dir(p.s.path), 0o755)
		if err != nil {
			return err
		}

		fi, err := os.Create(p.s.path)
		if err != nil {
			return err
		}

		p.s.f.a.agent.OnSegmentCreate(p.s.path)

		err = writeInit(fi, p.s.f.tracks)
		if err != nil {
			fi.Close()
			return err
		}

		p.s.fi = fi
	}

	return writePart(p.s.fi, p.sequenceNumber, p.partTracks)
}

func (p *formatFMP4Part) record(track *formatFMP4Track, sample *sample) error {
	partTrack, ok := p.partTracks[track]
	if !ok {
		partTrack = &fmp4.PartTrack{
			ID:       track.initTrack.ID,
			BaseTime: durationGoToMp4(sample.dts-p.s.startDTS, track.initTrack.TimeScale),
		}
		p.partTracks[track] = partTrack
	}

	partTrack.Samples = append(partTrack.Samples, sample.PartSample)
	p.endDTS = sample.dts

	return nil
}

func (p *formatFMP4Part) duration() time.Duration {
	return p.endDTS - p.startDTS
}
