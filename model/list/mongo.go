package list

import (
	"context"
	"fmt"

	"github.com/carlosescorche/qrlist/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var collection *mongo.Collection

func getCollection() (*mongo.Collection, error) {
	if collection != nil {
		return collection, nil
	}

	db := database.Database()

	collection = db.Collection("list")

	_, err := collection.Indexes().CreateOne(
		context.Background(),
		mongo.IndexModel{
			Keys: bson.M{
				"CreatedBy": 1,
			}, Options: nil,
		},
	)

	if err != nil {
		return nil, err
	}

	return collection, nil
}

func insert(l *List) error {
	collection, err := getCollection()
	if err != nil {
		return err
	}

	_, err = collection.InsertOne(context.Background(), l)
	if err != nil {
		return err
	}

	return nil
}

func findById(id string) (*List, error) {
	var collection, err = getCollection()
	if err != nil {
		return nil, err
	}

	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrListId, err)
	}

	list := &List{}
	filter := bson.M{"_id": _id}
	if err = collection.FindOne(context.Background(), filter).Decode(list); err != nil {
		return nil, err
	}

	return list, nil
}

func findSubsById(id string) (*List, error) {
	var collection, err = getCollection()
	if err != nil {
		return nil, err
	}

	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrListId, err)
	}

	list := &List{}
	filter := bson.M{"list": _id}
	if err = collection.FindOne(context.Background(), filter).Decode(list); err != nil {
		return nil, err
	}

	return list, nil
}
