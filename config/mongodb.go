package config

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func EnvMongoURI() string {
	if err := godotenv.Load(); err != nil {
		log.Fatal("ERROR: loading config from .env file")
	}

	return os.Getenv("MONGOURI")
}

func ConnectDB() *mongo.Client {
	// "mongodb://user:password@host:port/test?authSource=admin&replicaSet=Cluster0-shard-0&readPreference=primary&appname=MongoDB%20Compass&ssl=true"
	client, err := mongo.NewClient(options.Client().ApplyURI(EnvMongoURI()))
	if err != nil {
		log.Fatal("ERROR: ", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	err = client.Connect(ctx)

	defer cancel()

	err = client.Ping(context.Background(), readpref.Primary())

	if err != nil {
		log.Fatal("Could not connect to the database ", err)
	}

	log.Println("Connected!")

	return client
}
