package graph

import (
	"github.com/noranetworks/kvds/graph/model"
	"github.com/noranetworks/kvds/internal/core"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Core  *core.Core
	todos []*model.Todo
}
