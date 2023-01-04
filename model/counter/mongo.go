package counter

import (
	"context"
	"time"

	"github.com/carlosescorche/qrlist/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection

func getCollection() (*mongo.Collection, error) {
	if collection != nil {
		return collection, nil
	}

	db := database.Database()

	collection = db.Collection("counter")

	_, err := collection.Indexes().CreateMany(
		context.Background(),
		[]mongo.IndexModel{
			{
				Keys: bson.D{
					{Key: "listId", Value: 1},
				}, Options: nil,
			},
			{
				Keys: bson.D{
					{Key: "expireOn", Value: 1},
				}, Options: options.Index().SetExpireAfterSeconds(1),
			},
		},
	)

	if err != nil {
		return nil, err
	}

	return collection, nil
}

func getCounter(id primitive.ObjectID) (*Counter, error) {
	collection, err := getCollection()
	if err != nil {
		return nil, err
	}

	counter := &Counter{}
	filter := bson.M{"listId": id}
	if err := collection.FindOneAndUpdate(context.Background(), filter, bson.D{{"$setOnInsert", bson.D{{"createdAt", time.Now()}, {"expireOn", time.Now().Add(24 * time.Hour)}}}, {"$inc", bson.D{{"count", 1}}}}, options.FindOneAndUpdate().SetUpsert(true).SetReturnDocument(options.After)).Decode(counter); err != nil {
		return nil, err
	}

	return counter, nil
}
