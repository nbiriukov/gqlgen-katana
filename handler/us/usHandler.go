package handler

import (
	"github.com/go-chi/chi"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/raganmartinez-hf/gqlgen-katana/graphql/us/graph"
	"github.com/raganmartinez-hf/gqlgen-katana/graphql/us/graph/generated"
)

func InitUSHandler(router chi.Router) {
	usHandler := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	// US endpoints go here
	router.Handle("/us", playground.Handler("US GraphQL playground", "/query/us"))
	router.Handle("/query/us", usHandler)
}
