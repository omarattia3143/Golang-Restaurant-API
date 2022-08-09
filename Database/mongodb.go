package Database

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"
)

// Connection URI
const uri = "mongodb://0.0.0.0:27017"

var Client *mongo.Client

func Setup() {
	// Create a new client and connect to the server
	Client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatalln(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	err = Client.Connect(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	// Ping the primary
	if err := Client.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected and pinged.")
}

func OpenCollection(collectionName string) *mongo.Collection {
	return Client.Database("restaurant").Collection(collectionName)
}
