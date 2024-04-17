package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
)

func main() {
	// Set client options
	uri := os.Getenv("MONGO_URI")
	clientOptions := options.Client().ApplyURI(uri)

	// Connect to MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	// NOTE - important
	// ensure majority of nodes have written
	// add a 5s write timeout
	wc := writeconcern.Majority()
	wc.WTimeout = 5 * time.Second
	collOpts := options.Collection().SetWriteConcern(wc)
	collection := client.Database("testdb").Collection("numbers", collOpts)

	// Insert multiple documents
	num := 100
	for i := 1; i <= num; i++ {
		doc := bson.D{{Key: "id", Value: i}, {Key: "name", Value: fmt.Sprintf("Record %d", i)}}
		c := context.Background()
		_, err := collection.InsertOne(c, doc)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Inserted key #%d\n", i)
		time.Sleep(1000 * time.Millisecond)
	}

	fmt.Printf("Inserted %d records\n", num)

	// Close the connection once main function finishes
	err = client.Disconnect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connection to MongoDB closed.")
}
