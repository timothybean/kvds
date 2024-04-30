package core

import (
	"net/http"
	"time"

	// start pprof
	_ "net/http/pprof"

	"github.com/noranetworks/kvds/internal/conf"
	"github.com/noranetworks/kvds/internal/logger"
	"github.com/noranetworks/kvds/internal/protocols/httpserv"
	"github.com/noranetworks/kvds/internal/restrictnetwork"
)

type pprofParent interface {
	logger.Writer
}

type pprof struct {
	parent pprofParent

	httpServer *httpserv.WrappedServer
}

func newPPROF(
	address string,
	readTimeout conf.StringDuration,
	parent pprofParent,
) (*pprof, error) {
	pp := &pprof{
		parent: parent,
	}

	network, address := restrictnetwork.Restrict("tcp", address)

	var err error
	pp.httpServer, err = httpserv.NewWrappedServer(
		network,
		address,
		time.Duration(readTimeout),
		"",
		"",
		http.DefaultServeMux,
		pp,
	)
	if err != nil {
		return nil, err
	}

	pp.Log(logger.Info, "listener opened on "+address)

	return pp, nil
}

func (pp *pprof) close() {
	pp.Log(logger.Info, "listener is closing")
	pp.httpServer.Close()
}

func (pp *pprof) Log(level logger.Level, format string, args ...interface{}) {
	pp.parent.Log(level, "[pprof] "+format, args...)
}
