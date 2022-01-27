package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	_"fmt"

	"github.com/NutthakornS/todos/graph/model"
)

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	users := []*model.User{}
	user := model.User{
		ID:   "dang-001",
		Name: "Aj.Dang",
	}
	result := append(users, &user)
	return result, nil
}
