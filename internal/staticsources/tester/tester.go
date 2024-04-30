// Package tester contains a static source tester.
package tester

import (
	"context"

	"github.com/noranetworks/kvds/internal/asyncwriter"
	"github.com/noranetworks/kvds/internal/conf"
	"github.com/noranetworks/kvds/internal/defs"
	"github.com/noranetworks/kvds/internal/logger"
	"github.com/noranetworks/kvds/internal/stream"
	"github.com/noranetworks/kvds/internal/unit"
)

// Tester is a static source tester.
type Tester struct {
	ctx       context.Context
	ctxCancel func()
	stream    *stream.Stream
	writer    *asyncwriter.Writer

	Unit chan unit.Unit
	done chan struct{}
}

// New allocates a tester.
func New(createFunc func(defs.StaticSourceParent) defs.StaticSource, conf *conf.Path) *Tester {
	ctx, ctxCancel := context.WithCancel(context.Background())

	t := &Tester{
		ctx:       ctx,
		ctxCancel: ctxCancel,
		Unit:      make(chan unit.Unit),
		done:      make(chan struct{}),
	}

	s := createFunc(t)

	go func() {
		s.Run(defs.StaticSourceRunParams{ //nolint:errcheck
			Context: ctx,
			Conf:    conf,
		})
		close(t.done)
	}()

	return t
}

// Close closes the tester.
func (t *Tester) Close() {
	t.ctxCancel()
	t.writer.Stop()
	t.stream.Close()
	<-t.done
}

// Log implements StaticSourceParent.
func (t *Tester) Log(_ logger.Level, _ string, _ ...interface{}) {
}

// SetReady implements StaticSourceParent.
func (t *Tester) SetReady(req defs.PathSourceStaticSetReadyReq) defs.PathSourceStaticSetReadyRes {
	t.stream, _ = stream.New(
		1460,
		req.Desc,
		req.GenerateRTPPackets,
		t,
	)

	t.writer = asyncwriter.New(2048, t)
	t.stream.AddReader(t.writer, req.Desc.Medias[0], req.Desc.Medias[0].Formats[0], func(u unit.Unit) error {
		t.Unit <- u
		close(t.Unit)
		return nil
	})
	t.writer.Start()

	return defs.PathSourceStaticSetReadyRes{
		Stream: t.stream,
	}
}

// SetNotReady implements StaticSourceParent.
func (t *Tester) SetNotReady(_ defs.PathSourceStaticSetNotReadyReq) {
}
