package server

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/rokiyama/gqlgen-todos/graph"
	"github.com/rokiyama/gqlgen-todos/graph/generated"
)

func New() *handler.Server {
	return handler.NewDefaultServer(
		generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}),
	)
}
