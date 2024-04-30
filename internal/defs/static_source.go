package defs

import (
	"context"

	"github.com/noranetworks/kvds/internal/conf"
	"github.com/noranetworks/kvds/internal/logger"
)

// StaticSource is a static source.
type StaticSource interface {
	logger.Writer
	Run(StaticSourceRunParams) error
	APISourceDescribe() APIPathSourceOrReader
}

// StaticSourceParent is the parent of a static source.
type StaticSourceParent interface {
	logger.Writer
	SetReady(req PathSourceStaticSetReadyReq) PathSourceStaticSetReadyRes
	SetNotReady(req PathSourceStaticSetNotReadyReq)
}

// StaticSourceRunParams is the set of params passed to Run().
type StaticSourceRunParams struct {
	Context    context.Context
	Conf       *conf.Path
	ReloadConf chan *conf.Path
}
