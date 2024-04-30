package core

import (
	"context"
	"crypto/tls"
	"net"
	"sync"

	"github.com/noranetworks/kvds/internal/logger"
)

type snapshotServerParent interface {
	logger.Writer
}

type SnapshotServer struct {
	allowOrigin string
	parent      snapshotServerParent

	ctx       context.Context
	ctxCancel func()
	ln        net.Listener
	tlsConfig *tls.Config
	wg        *sync.WaitGroup

	done chan struct{}
}

func newSnapshotServer(
	parentCtx context.Context,
	address string,
	encryption bool,
	serverKey string,
	serverCert string,
	allowOrigin string,
	parent snapshotServerParent,
) (*SnapshotServer, error) {
	ln, err := net.Listen("tcp", address)
	if err != nil {
		return nil, err
	}

	var tlsConfig *tls.Config
	if encryption {
		crt, err := tls.LoadX509KeyPair(serverCert, serverKey)
		if err != nil {
			ln.Close()
			return nil, err
		}

		tlsConfig = &tls.Config{
			Certificates: []tls.Certificate{crt},
		}
	}

	ctx, ctxCancel := context.WithCancel(parentCtx)

	ss := &SnapshotServer{
		allowOrigin: allowOrigin,
		parent:      parent,
		ctx:         ctx,
		ctxCancel:   ctxCancel,
		ln:          ln,
		tlsConfig:   tlsConfig,

		done: make(chan struct{}),
	}

	ss.Log(logger.Debug, "Snapshot server created on: "+address)

	return ss, nil
}

// Log is the main logging function.
func (s *SnapshotServer) Log(level logger.Level, format string, args ...interface{}) {
	s.parent.Log(level, format, args...)
}
