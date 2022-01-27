package service

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/event"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	MongoClient *mongo.Client
	MongoDb     *mongo.Database
)

func ConnectToMongo() {
	var err error
	cmdMonitor := &event.CommandMonitor{
		Started: func(_ context.Context, evt *event.CommandStartedEvent) {
			log.Print(evt.Command)
		},
	}
	connectionString := "mongodb://root:admin123@localhost"
	clientOpts := options.Client().ApplyURI(connectionString).SetMonitor(cmdMonitor)
	MongoClient, err = mongo.Connect(context.TODO(), clientOpts)
	MongoDb = MongoClient.Database("pnpdb")

	if err != nil {
		panic(err)
	}
}
