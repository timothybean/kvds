package core

import (
	"context"
	"fmt"
	"sort"
	"sync"

	"github.com/noranetworks/kvds/internal/conf"
	"github.com/noranetworks/kvds/internal/defs"
	"github.com/noranetworks/kvds/internal/logger"
)

type hlsManagerAPIMuxersListRes struct {
	data *defs.APIHLSMuxerList
	err  error
}

type hlsManagerAPIMuxersListReq struct {
	res chan hlsManagerAPIMuxersListRes
}

type hlsManagerAPIMuxersGetRes struct {
	data *defs.APIHLSMuxer
	err  error
}

type hlsManagerAPIMuxersGetReq struct {
	name string
	res  chan hlsManagerAPIMuxersGetRes
}

type hlsManagerParent interface {
	logger.Writer
}

type hlsManager struct {
	externalAuthenticationURL string
	alwaysRemux               bool
	variant                   conf.HLSVariant
	segmentCount              int
	segmentDuration           conf.StringDuration
	partDuration              conf.StringDuration
	segmentMaxSize            conf.StringSize
	directory                 string
	writeQueueSize            int
	pathManager               *pathManager
	parent                    hlsManagerParent

	ctx        context.Context
	ctxCancel  func()
	wg         sync.WaitGroup
	httpServer *hlsHTTPServer
	muxers     map[string]*hlsMuxer

	// in
	chPathReady     chan *path
	chPathNotReady  chan *path
	chHandleRequest chan hlsMuxerHandleRequestReq
	chCloseMuxer    chan *hlsMuxer
	chAPIMuxerList  chan hlsManagerAPIMuxersListReq
	chAPIMuxerGet   chan hlsManagerAPIMuxersGetReq
}

func newHLSManager(
	address string,
	encryption bool,
	serverKey string,
	serverCert string,
	externalAuthenticationURL string,
	alwaysRemux bool,
	variant conf.HLSVariant,
	segmentCount int,
	segmentDuration conf.StringDuration,
	partDuration conf.StringDuration,
	segmentMaxSize conf.StringSize,
	allowOrigin string,
	trustedProxies conf.IPsOrCIDRs,
	directory string,
	readTimeout conf.StringDuration,
	writeQueueSize int,
	pathManager *pathManager,
	parent hlsManagerParent,
) (*hlsManager, error) {
	ctx, ctxCancel := context.WithCancel(context.Background())

	m := &hlsManager{
		externalAuthenticationURL: externalAuthenticationURL,
		alwaysRemux:               alwaysRemux,
		variant:                   variant,
		segmentCount:              segmentCount,
		segmentDuration:           segmentDuration,
		partDuration:              partDuration,
		segmentMaxSize:            segmentMaxSize,
		directory:                 directory,
		writeQueueSize:            writeQueueSize,
		pathManager:               pathManager,
		parent:                    parent,
		ctx:                       ctx,
		ctxCancel:                 ctxCancel,
		muxers:                    make(map[string]*hlsMuxer),
		chPathReady:               make(chan *path),
		chPathNotReady:            make(chan *path),
		chHandleRequest:           make(chan hlsMuxerHandleRequestReq),
		chCloseMuxer:              make(chan *hlsMuxer),
		chAPIMuxerList:            make(chan hlsManagerAPIMuxersListReq),
		chAPIMuxerGet:             make(chan hlsManagerAPIMuxersGetReq),
	}

	var err error
	m.httpServer, err = newHLSHTTPServer(
		address,
		encryption,
		serverKey,
		serverCert,
		allowOrigin,
		trustedProxies,
		readTimeout,
		m.pathManager,
		m,
	)
	if err != nil {
		ctxCancel()
		return nil, err
	}

	m.Log(logger.Info, "listener opened on "+address)

	m.wg.Add(1)
	go m.run()

	return m, nil
}

// Log is the main logging function.
func (m *hlsManager) Log(level logger.Level, format string, args ...interface{}) {
	m.parent.Log(level, "[HLS] "+format, args...)
}

func (m *hlsManager) close() {
	m.Log(logger.Info, "listener is closing")
	m.ctxCancel()
	m.wg.Wait()
}

func (m *hlsManager) run() {
	defer m.wg.Done()

outer:
	for {
		select {
		case pa := <-m.chPathReady:
			if m.alwaysRemux && !pa.conf.SourceOnDemand && pa.conf.Enabled {
				if _, ok := m.muxers[pa.name]; !ok {
					m.createMuxer(pa.name, "")
				}
			}

		case pa := <-m.chPathNotReady:
			c, ok := m.muxers[pa.name]
			if ok && c.remoteAddr == "" { // created with "always remux"
				c.close()
				delete(m.muxers, pa.name)
			}

		case req := <-m.chHandleRequest:
			source, err := m.pathManager.getSourceForPath(req.path, "HLS")
			if err != nil {
				fmt.Println(err)
			}

			r, ok := m.muxers[source]
			switch {
			case ok:
				r.processRequest(&req)

			default:
				r := m.createMuxer(req.path, req.ctx.ClientIP())
				r.processRequest(&req)
			}

		case c := <-m.chCloseMuxer:
			if c2, ok := m.muxers[c.PathName()]; !ok || c2 != c {
				continue
			}
			delete(m.muxers, c.PathName())

		case req := <-m.chAPIMuxerList:
			data := &defs.APIHLSMuxerList{
				Items: []*defs.APIHLSMuxer{},
			}

			for _, muxer := range m.muxers {
				data.Items = append(data.Items, muxer.apiItem())
			}

			sort.Slice(data.Items, func(i, j int) bool {
				return data.Items[i].Created.Before(data.Items[j].Created)
			})

			req.res <- hlsManagerAPIMuxersListRes{
				data: data,
			}

		case req := <-m.chAPIMuxerGet:
			muxer, ok := m.muxers[req.name]
			if !ok {
				req.res <- hlsManagerAPIMuxersGetRes{err: fmt.Errorf("muxer not found")}
				continue
			}

			req.res <- hlsManagerAPIMuxersGetRes{data: muxer.apiItem()}

		case <-m.ctx.Done():
			break outer
		}
	}

	m.ctxCancel()

	m.httpServer.close()
}

func (m *hlsManager) createMuxer(pathName string, remoteAddr string) *hlsMuxer {
	source, err := m.pathManager.getSourceForPath(pathName, "HLS")
	if err != nil {
		fmt.Println(err)
	}

	/* 	if pa, ok := m.pathManager.pathConfs[source]; ok {
	   		if !pa.Enabled {
	   			return nil
	   		}
	   	}
	*/
	r := newHLSMuxer(
		m.ctx,
		remoteAddr,
		m.externalAuthenticationURL,
		m.variant,
		m.segmentCount,
		m.segmentDuration,
		m.partDuration,
		m.segmentMaxSize,
		m.directory,
		m.writeQueueSize,
		&m.wg,
		source,
		m.pathManager,
		m)
	m.muxers[source] = r
	return r
}

// closeMuxer is called by hlsMuxer.
func (m *hlsManager) closeMuxer(c *hlsMuxer) {
	select {
	case m.chCloseMuxer <- c:
	case <-m.ctx.Done():
	}
}

// pathReady is called by pathManager.
func (m *hlsManager) pathReady(pa *path) {
	select {
	case m.chPathReady <- pa:
	case <-m.ctx.Done():
	}
}

// pathNotReady is called by pathManager.
func (m *hlsManager) pathNotReady(pa *path) {
	select {
	case m.chPathNotReady <- pa:
	case <-m.ctx.Done():
	}
}

// apiMuxersList is called by api.
func (m *hlsManager) apiMuxersList() (*defs.APIHLSMuxerList, error) {
	req := hlsManagerAPIMuxersListReq{
		res: make(chan hlsManagerAPIMuxersListRes),
	}

	select {
	case m.chAPIMuxerList <- req:
		res := <-req.res
		return res.data, res.err

	case <-m.ctx.Done():
		return nil, fmt.Errorf("terminated")
	}
}

// apiMuxersGet is called by api.
func (m *hlsManager) apiMuxersGet(name string) (*defs.APIHLSMuxer, error) {
	req := hlsManagerAPIMuxersGetReq{
		name: name,
		res:  make(chan hlsManagerAPIMuxersGetRes),
	}

	select {
	case m.chAPIMuxerGet <- req:
		res := <-req.res
		return res.data, res.err

	case <-m.ctx.Done():
		return nil, fmt.Errorf("terminated")
	}
}

func (m *hlsManager) handleRequest(req hlsMuxerHandleRequestReq) {
	req.res = make(chan *hlsMuxer)

	select {
	case m.chHandleRequest <- req:
		muxer := <-req.res
		if muxer != nil {
			req.ctx.Request.URL.Path = req.file
			muxer.handleRequest(req.ctx)
		}

	case <-m.ctx.Done():
	}
}
