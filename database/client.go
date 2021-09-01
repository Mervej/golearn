package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Connector *mongo.Collection
var ctx = context.TODO()

func Connect(connectionString string) error {
	var err error
	clientOptions := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		log.Println("err while connecting to db")
	}

	Connector = client.Database("vinnya").Collection("fixedImages")

	return nil
}
