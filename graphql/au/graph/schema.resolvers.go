package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/raganmartinez-hf/gqlgen-katana/graphql/au/graph/generated"
	"github.com/raganmartinez-hf/gqlgen-katana/graphql/au/graph/model"
)

func (r *mutationResolver) CreateHub(ctx context.Context, input model.NewHub) (*model.Hub, error) {
	hub := &model.Hub{
		ID:       "1",
		Country:  input.Country,
		Name:     input.Name,
		Address:  input.Address,
		State:    input.State,
		Priority: input.Priority,
	}
	r.hubs = append(r.hubs, hub)
	return hub, nil
}

func (r *queryResolver) Hubs(ctx context.Context) ([]*model.Hub, error) {
	return r.hubs, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
