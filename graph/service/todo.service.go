package service

import (
	"context"

	"github.com/NutthakornS/todos/graph/model"
	"go.mongodb.org/mongo-driver/bson"
)

type todoService struct {
	collection string
}

var (
	TodoService = todoService{collection: "todo"}
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

func (s *todoService) CreateTodo(input model.NewTodo) (*model.Todo, error) {
	col := MongoDb.Collection(s.collection)
	result, err := col.InsertOne(context.TODO(), &input)
	if err != nil {
		return nil, err
	} else {
		newU := model.Todo{}
		err := col.FindOne(context.TODO(), bson.M{"_id": result.InsertedID}).Decode(&newU)
		if err != nil {
			return nil, err
		}
		return &newU, nil
	}
}
