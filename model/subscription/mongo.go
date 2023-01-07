package subscription

import (
	"context"
	"fmt"

	"github.com/carlosescorche/qrlist-api/database"
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

	collection = db.Collection("subscriptions")

	_, err := collection.Indexes().CreateMany(
		context.Background(),
		[]mongo.IndexModel{{
			Keys:    bson.M{"userId": 1},
			Options: nil,
		}, {
			Keys:    bson.M{"listId": 1},
			Options: nil,
		}},
	)

	if err != nil {
		return nil, err
	}

	return collection, nil
}

func insert(s *Subscription) (*Subscription, error) {
	collection, err := getCollection()
	if err != nil {
		return nil, err
	}

	_, err = collection.InsertOne(context.Background(), s)
	if err != nil {
		return nil, err
	}

	return s, nil
}

func update(subs Subscription) error {
	var collection, err = getCollection()
	if err != nil {
		return err
	}

	filter := bson.M{"_id": subs.Id}
	if _, err := collection.UpdateOne(context.Background(), filter, bson.M{"$set": subs}); err != nil {
		return err
	}

	return nil
}

func findById(id string) (*Subscription, error) {
	var collection, err = getCollection()
	if err != nil {
		return nil, err
	}

	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrSubscriptionId, err)
	}

	subs := &Subscription{}
	filter := bson.M{"_id": _id}
	if err = collection.FindOne(context.Background(), filter).Decode(subs); err != nil {
		return nil, err
	}

	return subs, nil
}

func findByListId(id string) ([]Subscription, error) {
	var collection, err = getCollection()
	if err != nil {
		return nil, err
	}

	listId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrSubscriptionId, err)
	}

	filter := bson.M{"listId": listId}
	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}

	var subs []Subscription
	if err = cursor.All(context.TODO(), &subs); err != nil {
		return nil, err
	}

	return subs, nil
}
