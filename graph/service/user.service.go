package service

import "github.com/NutthakornS/todos/graph/model"

type userService struct{}

var (
	UserService = userService{}
)

func (s *userService) Users() ([]*model.User, error) {
	users := []*model.User{}
	usr := model.User{
		ID:   "snk-001",
		Name: "Warakorn",
	}
	result := append(users, &usr)
	return result, nil
}
