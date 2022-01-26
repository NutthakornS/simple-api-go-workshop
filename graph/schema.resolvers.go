package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	_"fmt"
	_ "os/user"

	"github.com/NutthakornS/todos/graph/generated"
	"github.com/NutthakornS/todos/graph/model"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	createdTodo := model.Todo{
		ID:   "yyyy-zzzz",
		Text: input.Text,
		Done: false,
		User: &model.User{
			ID:   input.UserID,
			Name: "Someone in the world",
		},
	}
	return &createdTodo, nil
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	result := []*model.Todo{}
	todo := model.Todo{
		ID:   "xxxx-yyyy",
		Text: "Do something",
		Done: true,
		User: &model.User{
			ID:   "user-001",
			Name: "Someone in the world",
		},
	}
	result = append(result, &todo)
	return result, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
