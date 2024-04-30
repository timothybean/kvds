package core

import (
	"context"
	"encoding/json"
	"fmt"
	"image"
	"sync"
	"time"

	"github.com/bluenviron/gortsplib/v4/pkg/description"
	"github.com/bluenviron/gortsplib/v4/pkg/format"
	"github.com/bluenviron/mediacommon/pkg/codecs/h264"
	"github.com/bluenviron/mediacommon/pkg/codecs/h265"
	"github.com/noranetworks/kvds/internal/asyncwriter"
	"github.com/noranetworks/kvds/internal/conf"
	"github.com/noranetworks/kvds/internal/defs"
	"github.com/noranetworks/kvds/internal/logger"
	"github.com/noranetworks/kvds/internal/stream"
	"github.com/noranetworks/kvds/internal/unit"
)

/* type snapshotPathManager interface {
	readerAdd(req pathReaderAddReq) pathReaderSetupPlayRes
} */

type snapshotParent interface {
	logger.Writer
}

/*
	 	Group is an output destination group
	 	name will be the url: {name}/stream
	 	stream is the output video name
		Enabled enabled or disables the group
		Defualt settings will be applied to the streams by default, but are configurable by steam
*/
type Snapshot struct {
	parent   snapshotParent
	Source   string `json:"source"`
	Enabled  bool   `json:"enabled"`
	Interval uint32 `json:"interval"`
	OnDemand bool   `json:"on_demand"`
	Width    uint16 `json:"width"`
	Height   uint16 `json:"height"`
	Image    image.Image
	//ReadBufferCount int

	Generate bool

	pathManager *pathManager
	path        *path
	stream      *stream.Stream
	videoMedia  *description.Media
	videoFormat format.Format
	//buffer       *bytes.Buffer
	//videoDecoder *SnapshotDecoder
	writer *asyncwriter.Writer

	wg        sync.WaitGroup
	ctx       context.Context
	ctxCancel func()

	chSourceReady    chan *path
	chSourceNotReady chan *path
	chConfReload     chan *conf.Snapshot

	done     chan struct{}
	generate chan bool
}

type SnapshotAddRes struct {
	Result bool `json:"result"`
	err    error
}

type SnapshotAddReq struct {
	/* SourceId string `json:"source_id"`
	Enabled  bool   `json:"enabled"`
	Interval uint8  `json:"interval"`
	OnDemand bool   `json:"on_demand"`
	Width    uint16 `json:"width"`
	Height   uint16 `json:"height"` */
	conf *conf.Snapshot
	res  chan SnapshotAddRes
}

type SnapshotRemoveRes struct {
	SourceId string `json:"source"`
	err      error
}

type SnapshotRemoveReq struct {
	SourceId string `json:"source_id"`
	res      chan SnapshotRemoveRes
}

type SnapshotUpdateRes struct {
	Snapshot Snapshot `json:"snapshot"`
	err      error
}

type SnapshotUpdateReq struct {
	Source   string `json:"source"`
	Enabled  bool   `json:"enabled"`
	Interval uint32 `json:"interval"`
	OnDemand bool   `json:"on_demand"`
	Width    uint16 `json:"width"`
	Height   uint16 `json:"height"`
	res      chan SnapshotUpdateRes
}

type SnapshotEnableRes struct {
	Snapshot Snapshot `json:"snapshot"`
	err      error
}

type SnapshotEnableReq struct {
	SnapshotId string `json:"snapshot_id"`
	Enabled    bool   `json:"enabled"`
	res        chan SnapshotEnableRes
}

type GetSnapshotRes struct {
	Snapshot Snapshot `json:"spapshot"`
	err      error
}

type GetSnapshotReq struct {
	SnapshotId string `json:"snapshot_id"`
	res        chan GetSnapshotRes
}

type GetSnapshotsRes struct {
	Snapshot []Snapshot `json:"snapsshots"`
	err      error
}

type GetSnapshotsReq struct {
	res chan GetSnapshotsRes
}

func newSnapshot(
	parentCtx context.Context,
	parent snapshotParent,
	snapConf *conf.Snapshot,
	//ReadBufferCount int,
	pathManager *pathManager,
) *Snapshot {
	ctx, ctxCancel := context.WithCancel(parentCtx)

	data, _ := json.Marshal(snapConf)

	s := &Snapshot{
		parent:    parent,
		ctx:       ctx,
		ctxCancel: ctxCancel,
		Source:    snapConf.Source,
		Enabled:   snapConf.Enabled,
		OnDemand:  snapConf.OnDemand,
		Interval:  snapConf.Interval,
		Width:     snapConf.Width,
		Height:    snapConf.Height,
		//ReadBufferCount: ReadBufferCount,
		pathManager: pathManager,

		chSourceReady:    make(chan *path),
		chSourceNotReady: make(chan *path),
		chConfReload:     make(chan *conf.Snapshot),

		done:     make(chan struct{}),
		generate: make(chan bool),
	}

	s.writer = asyncwriter.New(4096, s)
	s.writer.Start()

	json.Unmarshal(data, s)

	s.Log(logger.Debug, "Snapshot created: "+s.Source)

	s.wg.Add(1)
	go s.run()

	return s
}

// Log is the main logging function.
func (s *Snapshot) Log(level logger.Level, format string, args ...interface{}) {
	s.parent.Log(level, format, args...)
}

func (s *Snapshot) run() {
	defer close(s.done)
	defer s.wg.Done()

	innerCtx, innerCtxCancel := context.WithCancel(s.ctx)
	runErr := make(chan error)

outer:
	for {
		select {
		case snapConf := <-s.chConfReload:
			fmt.Println(snapConf)

		case pa := <-s.chSourceReady:
			if s.Enabled {
				s.Log(logger.Info, "Snaptshot source ready: %s", pa.name)
				res := s.pathManager.addReader(pathAddReaderReq{
					author: s,
					accessRequest: pathAccessRequest{
						name:     s.Source,
						skipAuth: true,
					},
				})
				if res.err != nil {
					s.Log(logger.Error, res.err.Error())
					continue
				}
				s.path = res.path
				s.stream = res.stream
				for _, media := range res.stream.Desc().Medias {
					if media.Type == "video" {
						s.videoMedia = media
						s.videoFormat = media.Formats[0]
					}
				}
				switch mediaDescription(s.videoMedia) {
				case "MPEG4Video":
					break
				case "H264":
					go func() {
						runErr <- s.snapshotProcessorH264(innerCtx)
					}()
				case "H265":
					go func() {
						runErr <- s.snapshotProcessorH265(innerCtx)
						s.wg.Add(1)
					}()
				}
				s.generate <- true

			}

		case <-s.chSourceNotReady:
			s.Log(logger.Info, "Stop generate Snap: %s", s.Source)
			s.generate <- false

		case snapConf := <-s.chConfReload:
			// stop the inner process
			innerCtxCancel()

			if snapConf.Source != s.Source {
				// restart the snapshot and attach to new source id
				s.path.removeReader(pathRemoveReaderReq{author: s})
				s.Source = snapConf.Source
			}

			s.Enabled = snapConf.Enabled
			s.Interval = snapConf.Interval
			s.OnDemand = snapConf.OnDemand
			s.Width = snapConf.Width
			s.Height = snapConf.Height

			// TODO figure out if path is running
			for _, path := range s.pathManager.paths {
				if path.confName == s.Source {
					if path.onDemandPublisherState == pathOnDemandStateReady {
						s.pathManager.doAddReader(pathAddReaderReq{
							author: s,
							accessRequest: pathAccessRequest{
								name:     s.Source,
								skipAuth: true,
							},
						})

						res := s.pathManager.describe(pathDescribeReq{accessRequest: pathAccessRequest{
							name:     s.Source,
							skipAuth: true,
						}})

						s.path = res.path
						s.stream = res.stream
						for _, media := range res.stream.Desc().Medias {
							if media.Type == "video" {
								s.videoMedia = media
								s.videoFormat = media.Formats[0]
							}
						}
						switch mediaDescription(s.videoMedia) {
						case "MPEG4Video":
							break
						case "H264":
							go func() {
								runErr <- s.snapshotProcessorH264(innerCtx)
								s.wg.Add(1)
							}()
						case "H265":
							go func() {
								runErr <- s.snapshotProcessorH265(innerCtx)
								s.wg.Add(1)
							}()
						}
						s.generate <- true
					}
				}
			}

		case err := <-runErr:
			s.Log(logger.Error, err.Error())
			innerCtxCancel()

		case <-s.ctx.Done():
			innerCtxCancel()
			break outer
		}
	}

	s.ctxCancel()
}

func (s *Snapshot) close() {
	s.Log(logger.Info, "snapshot process for "+s.Source+" shutting down")
	s.ctxCancel()
}

// apiReaderDescribe implements reader.
func (s *Snapshot) apiReaderDescribe() defs.APIPathSourceOrReader {
	return defs.APIPathSourceOrReader{
		Type: "snapshot",
		ID:   s.Source,
	}
}

func (s *Snapshot) snapshotProcessorH264(ctx context.Context) error {
	s.Log(logger.Debug, "Running H264 Snapshot generator: %s", s.path.conf.Name)
	var mutex sync.Mutex

	defer func() {
		s.stream.RemoveReader(s.writer)
		s.path.removeReader(pathRemoveReaderReq{author: s})
	}()

	var videoFormat *format.H264
	s.videoMedia.FindFormat(&videoFormat)
	if videoFormat == nil {
		return fmt.Errorf("Track does not contain H.264 track")
	}

	rtpDecoder, err := videoFormat.CreateDecoder()
	if err != nil {
		return fmt.Errorf("Error creating H.264 packet decoder")
	}

	decoder, err := NewSnapshotDecoder("H264")
	if err != nil {
		return fmt.Errorf("Error creating H.264 decoder: " + err.Error())
	}
	defer func() {
		mutex.Lock()
		decoder.close()
		mutex.Unlock()
	}()

	sps, pps := videoFormat.SafeParams()
	if sps != nil {
		decoder.decodeSnapshotFrame(sps, 0, 0)
	}

	if pps != nil {
		decoder.decodeSnapshotFrame(pps, 0, 0)
	}

	ticker := time.NewTicker(time.Duration(float64(time.Second) * float64(s.Interval)))
	defer ticker.Stop()

	generateSnaps := true

	if s.OnDemand {
		ticker.Stop()
		generateSnaps = false
	}

	s.stream.AddReader(s.writer, s.videoMedia, s.videoFormat, func(dat unit.Unit) error {
		packets := dat.GetRTPPackets()

		for _, pkt := range packets {
			nalu, err := rtpDecoder.Decode(pkt)
			if err != nil {
				s.Log(logger.Debug, err.Error())
				continue
			}

			if generateSnaps && h264.IDRPresent(nalu) {
				for _, buffer := range nalu {

					mutex.Lock()
					img, err := decoder.decodeSnapshotFrame(buffer, int(s.Width), int(s.Height))
					mutex.Unlock()

					if err != nil {
						s.Log(logger.Error, "ffmpeg error: "+err.Error())
						continue
					}

					if img != nil {
						s.Image = img
						//fmt.Println("generate snap " + s.Source)

						if s.Enabled && !s.OnDemand {
							//go imaging.Save(s.Image, s.path.name+".jpg")
							generateSnaps = false
						}
					}
				}
			}
		}

		return nil
	})

outer:
	for {
		select {
		case gen := <-s.generate:
			generateSnaps = gen
		case <-ticker.C:
			generateSnaps = true

		case <-ctx.Done():
			ticker.Stop()
			generateSnaps = false
			break outer
		}
	}

	return nil
}

func (s *Snapshot) snapshotProcessorH265(ctx context.Context) error {
	s.Log(logger.Debug, "Running H265 Snapshot generator: %s", s.path.conf.Name)
	var mutex sync.Mutex

	defer func() {
		s.stream.RemoveReader(s.writer)
		s.path.removeReader(pathRemoveReaderReq{author: s})
	}()

	var videoFormat *format.H265
	s.videoMedia.FindFormat(&videoFormat)
	if videoFormat == nil {
		return fmt.Errorf("Track does not contain H.265 track")
	}

	rtpDecoder, err := videoFormat.CreateDecoder()
	if err != nil {
		return fmt.Errorf("Error creating H.265 packet decoder")
	}

	decoder, err := NewSnapshotDecoder("H265")
	if err != nil {
		return fmt.Errorf("Error creating H.265 decoder: " + err.Error())
	}
	defer func() {
		mutex.Lock()
		decoder.close()
		mutex.Unlock()
	}()

	vps, sps, pps := videoFormat.SafeParams()

	if vps != nil {
		decoder.decodeSnapshotFrame(vps, 0, 0)
	}

	if sps != nil {
		decoder.decodeSnapshotFrame(sps, 0, 0)
	}

	if pps != nil {
		decoder.decodeSnapshotFrame(pps, 0, 0)
	}

	ticker := time.NewTicker(time.Duration(float64(time.Second) * float64(s.Interval)))
	defer ticker.Stop()

	generateSnaps := true

	if s.OnDemand {
		ticker.Stop()
		generateSnaps = false
	}

	s.stream.AddReader(s.writer, s.videoMedia, s.videoFormat, func(dat unit.Unit) error {
		packets := dat.GetRTPPackets()

		for _, pkt := range packets {
			nalu, err := rtpDecoder.Decode(pkt)
			if err != nil {
				s.Log(logger.Debug, err.Error())
				continue
			}

			if generateSnaps && h265.IsRandomAccess(nalu) {
				for _, buffer := range nalu {

					mutex.Lock()
					img, err := decoder.decodeSnapshotFrame(buffer, int(s.Width), int(s.Height))
					mutex.Unlock()

					if err != nil {
						s.Log(logger.Error, "ffmpeg error: "+err.Error())
						continue
					}

					if img != nil {
						s.Image = img
						//fmt.Println("generate snap " + s.Source)

						if s.Enabled && !s.OnDemand {
							//go imaging.Save(s.Image, s.path.name+".jpg")
							generateSnaps = false
						}
					}
				}
			}
		}

		return nil
	})

outer:
	for {
		select {
		case gen := <-s.generate:
			generateSnaps = gen
		case <-ticker.C:
			generateSnaps = true

		case <-ctx.Done():
			ticker.Stop()
			generateSnaps = false
			break outer
		}
	}

	return nil

}
