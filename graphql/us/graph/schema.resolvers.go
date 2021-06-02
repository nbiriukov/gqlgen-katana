package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/raganmartinez-hf/gqlgen-katana/graphql/us/graph/generated"
	"github.com/raganmartinez-hf/gqlgen-katana/graphql/us/graph/model"
	"github.com/raganmartinez-hf/gqlgen-katana/utils/json"
)

func (r *mutationResolver) CreateHub(ctx context.Context, input model.NewHub) (*model.Hub, error) {
	hub := &model.Hub{
		ID:      "1",
		Country: input.Country,
		Name:    input.Name,
		Address: input.Address,
		State:   input.State,
	}
	r.hubs = append(r.hubs, hub)
	return hub, nil
}

func (r *queryResolver) Hubs(ctx context.Context) ([]*model.Hub, error) {
	return r.hubs, nil
}

func (r *queryResolver) Schema(ctx context.Context, id *string) (interface{}, error) {
	path := fmt.Sprintf("graphql/us/schemas/%s.json", *id)
	return json.ImportFromFile(path)
}

func (r *queryResolver) Menu(ctx context.Context) (interface{}, error) {
	return json.ImportFromFile("graphql/us/schemas/menu.json")
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
