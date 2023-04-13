package config

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func Connect() {
	clientOptions := options.Client().ApplyURI("mongo_url")
	client, err := mongo.NewClient(clientOptions)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	err = client.Connect(ctx)

	defer cancel()

	err = client.Ping(context.Background(), readpref.Primary())

	if err != nil {
		log.Fatal("Could not connect to the database ", err)
	} else {
		log.Println("Connected!")
	}

	// db := client.Database("barber_schedule")

	return
}
