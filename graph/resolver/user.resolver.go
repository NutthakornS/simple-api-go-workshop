package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/NutthakornS/todos/graph/model"
	"github.com/NutthakornS/todos/graph/service"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	return service.UserService.CreateUser(input)
}

func (r *mutationResolver) UpdateUser(ctx context.Context, id string, body model.NewUser) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	return service.UserService.Users()
}

func (r *queryResolver) FindUserByName(ctx context.Context, name string) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}
