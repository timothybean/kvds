package core

import (
	"context"
	"fmt"
	"image"
	"sync"

	"github.com/noranetworks/kvds/internal/conf"
	"github.com/noranetworks/kvds/internal/defs"
	"github.com/noranetworks/kvds/internal/logger"
)

type snapshotManagerParent interface {
	logger.Writer
}

type SnapshotManager struct {
	Snaptshots  map[string]*Snapshot
	pathManager *pathManager
	metrics     *metrics
	parent      snapshotManagerParent

	snapConfs map[string]*conf.Snapshot

	ctx       context.Context
	ctxCancel func()
	wg        sync.WaitGroup

	chConfReload chan map[string]*conf.Snapshot

	chAddSnapshot      chan SnapshotAddReq
	chUpdateSnapshot   chan SnapshotUpdateReq
	chRemoveSnapshot   chan SnapshotRemoveReq
	chEnableSnapshot   chan SnapshotEnableReq
	chGetSnapshot      chan GetSnapshotReq
	chGetSnapshots     chan GetSnapshotsReq
	chAPISnapshotsGet  chan snapshotAPISnapshotsGetReq
	chAPISnapshotsList chan snapshotAPISnapshotsListReq

	chPathSourceReady    chan *path
	chPathSourceNotReady chan *path
}

type snapshotAPISnapshotsListRes struct {
	Snapshots map[string]*conf.Snapshot
	err       error
}

type snapshotAPISnapshotsListReq struct {
	res chan snapshotAPISnapshotsListRes
}

type snapshotAPISnapshotsGetRes struct {
	Image image.Image
	err   error
}

type snapshotAPISnapshotsGetReq struct {
	name string
	res  chan snapshotAPISnapshotsGetRes
}

func newSnapshotManager(
	parentCtx context.Context,
	snapConfs map[string]*conf.Snapshot,
	pathManager *pathManager,
	metrics *metrics,
	parent snapshotManagerParent,
) *SnapshotManager {
	ctx, ctxCancel := context.WithCancel(parentCtx)

	sm := &SnapshotManager{
		snapConfs:   snapConfs,
		metrics:     metrics,
		parent:      parent,
		pathManager: pathManager,
		ctx:         ctx,
		ctxCancel:   ctxCancel,

		Snaptshots:         make(map[string]*Snapshot),
		chConfReload:       make(chan map[string]*conf.Snapshot),
		chAddSnapshot:      make(chan SnapshotAddReq),
		chUpdateSnapshot:   make(chan SnapshotUpdateReq),
		chRemoveSnapshot:   make(chan SnapshotRemoveReq),
		chEnableSnapshot:   make(chan SnapshotEnableReq),
		chGetSnapshot:      make(chan GetSnapshotReq),
		chGetSnapshots:     make(chan GetSnapshotsReq),
		chAPISnapshotsGet:  make(chan snapshotAPISnapshotsGetReq),
		chAPISnapshotsList: make(chan snapshotAPISnapshotsListReq),

		chPathSourceReady:    make(chan *path),
		chPathSourceNotReady: make(chan *path),
	}

	// Create groups from config
	for snapConfName, snapConf := range sm.snapConfs {
		sm.createSnapshot(snapConfName, snapConf, snapConfName, nil)
	}

	/* fmt.Println(sm.Snaptshots) */

	sm.pathManager.setSnapshotManager(sm)
	sm.Log(logger.Debug, "Snapshot Manager Created")

	sm.wg.Add(1)
	go sm.run()

	return sm
}

// Log is the main logging function.
func (sm *SnapshotManager) Log(level logger.Level, format string, args ...interface{}) {
	sm.parent.Log(level, "[%s] "+format, append([]interface{}{"SM"}, args...)...)
}

func (sm *SnapshotManager) close() {
	sm.Log(logger.Info, "Snapshot Manager closing")
	sm.ctxCancel()
	sm.wg.Wait()
}

func (sm *SnapshotManager) run() {
	defer sm.wg.Done()

outer:
	for {
		//fmt.Println("Snaps running")
		select {
		/* case conf := <-gm.chConfReload:
		 */
		case snap := <-sm.chAddSnapshot:
			fmt.Printf("[hangFunction] Looping 1 \n")
			if _, ok := sm.Snaptshots[snap.conf.Source]; !ok {
				sm.createSnapshot(snap.conf.Source, snap.conf, snap.conf.Source, nil)
				snap.res <- SnapshotAddRes{Result: true, err: nil}
				continue
			}

			snap.res <- SnapshotAddRes{Result: false, err: fmt.Errorf("Snapshot for " + snap.conf.Source + " already exists.")}

		case snap := <-sm.chRemoveSnapshot:
			fmt.Printf("[hangFunction] Looping 2 \n")
			if s, ok := sm.Snaptshots[snap.SourceId]; ok {
				s.ctxCancel()
				delete(sm.Snaptshots, snap.SourceId)
				snap.res <- SnapshotRemoveRes{SourceId: s.Source, err: nil}
				continue
			}

			snap.res <- SnapshotRemoveRes{SourceId: snap.SourceId, err: fmt.Errorf("Snapshot source id does not exist.")}

		case snap := <-sm.chUpdateSnapshot:
			fmt.Printf("[hangFunction] Looping 3 \n")
			if s, ok := sm.Snaptshots[snap.Source]; ok {
				s.chConfReload <- &conf.Snapshot{
					Source:   snap.Source,
					Enabled:  snap.Enabled,
					Interval: snap.Interval,
					OnDemand: snap.OnDemand,
					Width:    snap.Width,
					Height:   snap.Height,
				}
			}

		case source := <-sm.chPathSourceReady:
			fmt.Printf("[hangFunction] Looping 4 \n")
			if s, ok := sm.Snaptshots[source.confName]; ok {
				s.chSourceReady <- source
			}

		case source := <-sm.chPathSourceNotReady:
			fmt.Printf("[hangFunction] Looping 5 \n")
			if s, ok := sm.Snaptshots[source.confName]; ok {
				s.generate <- false
			}

		case req := <-sm.chAPISnapshotsList:
			fmt.Printf("[hangFunction] Looping 6 \n")
			req.res <- snapshotAPISnapshotsListRes{Snapshots: sm.snapConfs, err: nil}

		case req := <-sm.chAPISnapshotsGet:
			fmt.Printf("[hangFunction] Looping 7 \n")
			var snap *Snapshot
			if s, ok := sm.Snaptshots[req.name]; ok {
				sm.Log(logger.Error, "Snapshot %s not found.", req.name)
				if s.Image != nil {
					snap = s
				}
			}

			/* for _, s := range sm.Snaptshots {
				if s.path.conf.Name == req.name {
					if s.Image != nil {
						snap = s
					} else {
						sm.Log(logger.Error, "Snapshot %s image was null.", req.name)
					}
				}
			} */

			if snap != nil {
				req.res <- snapshotAPISnapshotsGetRes{Image: snap.Image, err: nil}
				continue
			}

			req.res <- snapshotAPISnapshotsGetRes{Image: nil, err: fmt.Errorf("No snapshot found for: " + req.name)}

		case <-sm.ctx.Done():
			fmt.Println("snaps break")
			break outer
		}

	}

	sm.ctxCancel()
}

func (sm *SnapshotManager) createSnapshot(snapConfName string, snapConf *conf.Snapshot, name string, matches []string) {
	s := newSnapshot(sm.ctx, sm, snapConf, sm.pathManager)
	sm.Snaptshots[snapConf.Source] = s
}

// pathSourceReady is called by pathManager.
func (s *SnapshotManager) pathReady(pa *path) {
	select {
	case s.chPathSourceReady <- pa:
	case <-s.ctx.Done():
	}
}

// pathSourceNotReady is called by pathManager.
func (s *SnapshotManager) pathNotReady(pa *path) {
	select {
	case s.chPathSourceNotReady <- pa:
	case <-s.ctx.Done():
	}
}

// apiSnapshotsList is called by api.
func (sm *SnapshotManager) apiSnaphotsList() (*defs.APISnapshotList, error) {
	req := snapshotAPISnapshotsListReq{
		res: make(chan snapshotAPISnapshotsListRes),
	}

	select {
	case sm.chAPISnapshotsList <- req:
		res := <-req.res

		return &defs.APISnapshotList{Snapshots: res.Snapshots, Err: res.err}, res.err

	case <-sm.ctx.Done():
		return nil, fmt.Errorf("terminated")
	}
}

// apiSnapshotsGet is called by api.
func (sm *SnapshotManager) apiSnapshotsGet(name string) (*defs.APISnapshotsGet, error) {
	req := snapshotAPISnapshotsGetReq{
		name: name,
		res:  make(chan snapshotAPISnapshotsGetRes),
	}

	select {
	case sm.chAPISnapshotsGet <- req:
		res := <-req.res

		return &defs.APISnapshotsGet{Image: res.Image, Err: res.err}, res.err

	case <-sm.ctx.Done():
		return nil, fmt.Errorf("terminated")
	}
}
