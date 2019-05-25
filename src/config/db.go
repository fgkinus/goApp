package config

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient = mongo.Client{}

func ConnectToDb(dbUri string) error {
	clientOptions := options.Client().ApplyURI(dbUri)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return err
	}
	// check the connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		return err
	} else {
		MongoClient = *client
		Logger.Info("Connected to mongo database on " + dbUri)
		return nil
	}
}
