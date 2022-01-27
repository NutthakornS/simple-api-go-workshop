package service

import "github.com/NutthakornS/todos/graph/model"

type todoService struct{}

var (
	TodoService = todoService{}
)

func (s *todoService) Todos() ([]*model.Todo, error) {
	result := []*model.Todo{}
	user := model.User{
		ID:   "milos-123",
		Name: "Riccado",
		// LastName: "Milos",
	}
	todo := model.Todo{
		ID:   "todo-milos-123",
		Text: "Dance practice",
		Done: true,
		User: &user,
	}
	res := append(result, &todo)
	return res, nil
}
