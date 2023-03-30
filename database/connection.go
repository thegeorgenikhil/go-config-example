package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect(uri string) *mongo.Client {
	options := options.Client().ApplyURI(uri)

	client, err := mongo.NewClient(options)

	if err != nil {
		log.Panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	err = client.Connect(ctx)

	if err != nil {
		log.Panic(err)
	}

	// Check the connection
	err = client.Ping(ctx, nil)

	if err != nil {
		log.Panic(err)
	}

	fmt.Println("Connected to MongoDB")

	return client
}
