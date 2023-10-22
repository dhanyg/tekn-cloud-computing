package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func connectToMongoDB(uri string) (*mongo.Client, context.Context, context.CancelFunc, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	return client, ctx, cancel, err
}

func getCollection(client *mongo.Client, dbName, collectionName string) (*mongo.Collection, error) {
	database := client.Database(dbName)
	collection := database.Collection(collectionName)
	return collection, nil
}

func fetchData(collection *mongo.Collection, ctx context.Context, filter bson.M) ([]bson.M, error) {
	var results []bson.M
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var result bson.M
		if err := cursor.Decode(&result); err != nil {
			return nil, err
		}
		results = append(results, result)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return results, nil
}

func main() {
	uri := "mongodb+srv://username:password@cluster0.vrbzjyd.mongodb.net/?retryWrites=true&w=majority"
	dbName := "tekn_cloud"
	collectionName := "persons"
	filter := bson.M{}

	client, ctx, cancel, err := connectToMongoDB(uri)
	if err != nil {
		log.Fatal(err)
		cancel()
	}
	defer client.Disconnect(ctx)
	fmt.Println("Database connected...")

	collection, err := getCollection(client, dbName, collectionName)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Get all persons data:")
	results, err := fetchData(collection, ctx, filter)
	if err != nil {
		log.Fatal(err)
	}

	for _, result := range results {
		fmt.Println(result)
	}

}
