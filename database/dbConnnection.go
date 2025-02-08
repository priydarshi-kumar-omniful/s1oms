package database

import (
	"context"
	"fmt"
	"log"
	"time"

	
	"oms/constants"


	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var Client *mongo.Client
var Ctx context.Context

func Connect(ctx context.Context) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	clientOptions := options.Client().ApplyURI(constants.MongoDBURI)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	Client = client
	Ctx = ctx

	// Check connection with a timeout
	err = Client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal("Could not connect to MongoDB:", err)
	} else {
		log.Println("Connected to MongoDB!")
	}
}

func GetCollection(collectionName string) *mongo.Collection {
	// Use context with timeout for MongoDB queries
	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	fmt.Println("get collection called")
	// Return the MongoDB collection with the context
	return Client.Database("oms_Service").Collection(collectionName)

}
