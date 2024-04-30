package record

import (
	"os"
	"path/filepath"
	"time"

	"github.com/noranetworks/kvds/internal/logger"
)

type formatMPEGTSSegment struct {
	f         *formatMPEGTS
	startDTS  time.Duration
	lastFlush time.Duration

	created time.Time
	path    string
	fi      *os.File
}

func newFormatMPEGTSSegment(f *formatMPEGTS, startDTS time.Duration) *formatMPEGTSSegment {
	s := &formatMPEGTSSegment{
		f:         f,
		startDTS:  startDTS,
		lastFlush: startDTS,
		created:   timeNow(),
	}

	f.dw.setTarget(s)

	return s
}

func (s *formatMPEGTSSegment) close() error {
	err := s.f.bw.Flush()

	if s.fi != nil {
		s.f.a.agent.Log(logger.Debug, "closing segment %s", s.path)
		err2 := s.fi.Close()
		if err == nil {
			err = err2
		}

		if err2 == nil {
			s.f.a.agent.OnSegmentComplete(s.path)
		}
	}

	return err
}

func (s *formatMPEGTSSegment) Write(p []byte) (int, error) {
	if s.fi == nil {
		s.path = path(s.created).encode(s.f.a.pathFormat)
		s.f.a.agent.Log(logger.Debug, "creating segment %s", s.path)

		err := os.MkdirAll(filepath.Dir(s.path), 0o755)
		if err != nil {
			return 0, err
		}

		fi, err := os.Create(s.path)
		if err != nil {
			return 0, err
		}

		s.f.a.agent.OnSegmentCreate(s.path)

		s.fi = fi
	}

	return s.fi.Write(p)
}
