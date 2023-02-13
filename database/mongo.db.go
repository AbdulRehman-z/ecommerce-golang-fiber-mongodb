package database

import (
	"context"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

func ConnectWithMongodb() *mongo.Client {

	// load .env file
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	// set mongodb connection string
	uri := os.Getenv("MONGO_URI")
	if uri == "" {
		log.Fatal("MONGODB_URI is not set")
	}
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}

	// check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Connected to mongodb")
	}

	return client
}

var Client *mongo.Client = ConnectWithMongodb()

// OpenCollection get collection
func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	var collection *mongo.Collection = client.Database("candy").Collection(collectionName)
	return collection
}
