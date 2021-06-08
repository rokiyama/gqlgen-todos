package graph

import (
	"github.com/graph-gophers/dataloader"
	"github.com/rokiyama/gqlgen-todos/graph/model"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Loader *dataloader.Loader
	todos  []*model.Todo
}
