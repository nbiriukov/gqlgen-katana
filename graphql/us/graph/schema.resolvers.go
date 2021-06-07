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
	r.hubAI++

	hub := &model.Hub{
		ID:      r.hubAI,
		Country: input.Country,
		Name:    input.Name,
		Address: input.Address,
		State:   input.State,
	}
	r.hubs = append(r.hubs, hub)
	return hub, nil
}

func (r *mutationResolver) UpdateHub(ctx context.Context, id int, input model.NewHub) (*model.Hub, error) {
	for _, hub := range r.hubs {
		if *(&hub.ID) == id {
			hub.Country = input.Country
			hub.Name = input.Name
			hub.Address = input.Address
			hub.State = input.State
		}

		return hub, nil
	}

	return nil, fmt.Errorf("not found")
}

func (r *mutationResolver) DeleteHub(ctx context.Context, id int) (*bool, error) {
	filtered := make([]*model.Hub, 0)
	for _, hub := range r.hubs {
		if *(&hub.ID) != id {
			filtered = append(filtered, hub)
		}
	}
	r.hubs = filtered

	return nil, nil
}

func (r *queryResolver) Hub(ctx context.Context, id *int) ([]*model.Hub, error) {
	if id == nil {
		return r.hubs, nil
	}

	filtered := make([]*model.Hub, 0)
	for _, hub := range r.hubs {
		if *(&hub.ID) == *id {
			filtered = append(filtered, hub)
		}
	}

	return filtered, nil
}

func (r *queryResolver) Schema(ctx context.Context) (interface{}, error) {
	return json.ImportFromFile("graphql/us/schema.json")
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
