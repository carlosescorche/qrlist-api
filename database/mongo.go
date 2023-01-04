package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoClient struct {
	ctx      context.Context
	client   *mongo.Client
	database string
}

var c *MongoClient

func Connect(uri string, database string) {
	ctx := context.TODO()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		panic(err)
	}

	log.Println("Connected to MongoDB")

	c = &MongoClient{
		ctx:      ctx,
		client:   client,
		database: database,
	}
}

func Close() {
	log.Println("Disconneting MongoDB")

	if err := c.client.Disconnect(c.ctx); err != nil {
		log.Fatal("Error disconnecting MongoDB: ", err)
	}
}

func Database() *mongo.Database {
	return c.client.Database(c.database)
}
