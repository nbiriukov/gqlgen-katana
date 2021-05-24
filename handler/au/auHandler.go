package handler

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/raganmartinez-hf/gqlgen-katana/graphql/au/graph"
	"github.com/raganmartinez-hf/gqlgen-katana/graphql/au/graph/generated"
)

func InitAUHandler() {
	auHandler := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	// AU endpoints go here
	http.Handle("/au", playground.Handler("AU GraphQL playground", "/query/au"))
	http.Handle("/query/au", auHandler)
}
