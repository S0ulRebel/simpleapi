package database

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MGDB *mongo.Client

func ConnectToMongoDB() (*mongo.Client, error) {
	var uri string
	if os.Getenv("MONGO_PASSWORD") != "" {
		uri = fmt.Sprintf(
			"mongodb://%s:%s@%s:%s/%s",
			os.Getenv("MONGO_USER"),
			os.Getenv("MONGO_PASSWORD"),
			os.Getenv("MONGO_HOST"),
			os.Getenv("MONGO_PORT"),
			os.Getenv("MONGO_DB"))
	} else {
		uri = fmt.Sprintf(
			"mongodb://%s:%s/%s",
			os.Getenv("MONGO_HOST"),
			os.Getenv("MONGO_PORT"),
			os.Getenv("MONGO_DB"))
	}

	client, err := mongo.Connect(
		context.TODO(),
		options.Client().ApplyURI(uri))

	if err != nil {
		return nil, err
	}

	MGDB = client
	return client, nil
}
