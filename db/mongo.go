package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	Client     *mongo.Client
	Database   *mongo.Database
	Collection *mongo.Collection
)

// InitMongoDB initializes MongoDB connection.
func InitMongoDB() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	Client = client
	Database = client.Database("advancedtodo")
	Collection = Database.Collection("advancedtodo")
}
