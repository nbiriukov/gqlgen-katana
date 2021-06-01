package handler

import (
	"github.com/go-chi/chi"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/raganmartinez-hf/gqlgen-katana/graphql/au/graph"
	"github.com/raganmartinez-hf/gqlgen-katana/graphql/au/graph/generated"
)

func InitAUHandler(router chi.Router) {
	auHandler := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	// AU endpoints go here
	router.Handle("/au", playground.Handler("AU GraphQL playground", "/query/au"))
	router.Handle("/query/au", auHandler)
}
