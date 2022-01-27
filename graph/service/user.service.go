package service

import (
	"context"

	"github.com/NutthakornS/todos/graph/model"
	"go.mongodb.org/mongo-driver/bson"
)

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

func (s *userService) CreateUser(input model.NewUser) (*model.User, error) {
	col := MongoDb.Collection("user")
	result, err := col.InsertOne(context.TODO(), input)
	if err != nil {
		return nil, err
	} else {
		newU := model.User{}
		err := col.FindOne(context.TODO(), bson.M{"_id": result.InsertedID}).Decode(&newU)
		if err != nil {
			return nil, err
		}
		return &newU, nil
	}
}
