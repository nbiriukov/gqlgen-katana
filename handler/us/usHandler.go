package handler

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/raganmartinez-hf/gqlgen-katana/graphql/us/graph"
	"github.com/raganmartinez-hf/gqlgen-katana/graphql/us/graph/generated"
)

func InitUSHandler() {
	usHandler := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	// US endpoints go here
	http.Handle("/us", playground.Handler("US GraphQL playground", "/query/us"))
	http.Handle("/query/us", usHandler)
}
