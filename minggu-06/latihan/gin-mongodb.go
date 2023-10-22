package main

import (
	"context"
	// "fmt"
	"log"
	"net/http"
	// "time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection

func connectToMongoDB(uri string) (*mongo.Client, context.Context, error) {
	ctx := context.TODO()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	return client, ctx, err
}

func fetchData(ctx context.Context, filter bson.M) ([]bson.M, error) {
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

func insertData(ctx context.Context, data interface{}) (*mongo.InsertOneResult, error) {
	result, err := collection.InsertOne(ctx, data)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func main() {
	uri := "mongodb+srv://username:password@cluster0.vrbzjyd.mongodb.net/?retryWrites=true&w=majority"
	dbName := "tekn_cloud"
	collectionName := "persons"
	filter := bson.M{}

	client, ctx, err := connectToMongoDB(uri)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	collection = client.Database(dbName).Collection(collectionName)

	r := gin.Default()

	r.GET("/persons", func(c *gin.Context) {
		results, err := fetchData(ctx, filter)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, results)
	})

	r.POST("/persons", func(c *gin.Context) {
		var data bson.M
		if err := c.BindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		insertResult, err := insertData(ctx, data)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
	}

	c.JSON(http.StatusCreated, gin.H{"insertedID": insertResult.InsertedID})
	})

	r.Run(":8080")
}
