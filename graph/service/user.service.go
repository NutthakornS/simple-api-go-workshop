package service

import (
	"context"

	"github.com/NutthakornS/todos/graph/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type userService struct{}

var (
	UserService = userService{}
)

func (s *userService) Users() ([]*model.User, error) {
	col := MongoDb.Collection("user")
	cursor, err := col.Find(context.TODO(), bson.M{}, &opts)
	var users []*model.User
	cursor.All(context.TODO(), &users)
	return users, err
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

func (s *userService) DeleteUser(id string) (bool, error) {
	col := MongoDb.Collection("user")
	nId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return false, err
	}
	res := col.FindOneAndDelete(context.TODO(), bson.M{"_id": nId})
	if res.Err() != nil {
		return false, res.Err()
	} else {
		return true, nil
	}
}
