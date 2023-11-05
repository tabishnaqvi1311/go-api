package config

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func connectMongo()(*mongo.Client, error){
	uri := os.Getenv("MONGO_URI")
	
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	return client, err
}

func getCollection(client *mongo.Client, name string)*mongo.Collection{
	coll := client.Database("test").Collection("books")
	return coll
}