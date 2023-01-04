package subscription

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	PENDING  = "PENDING"
	ACEPPTED = "ACCEPTED"
	FINISHED = "FINISHED"
)

type Subscription struct {
	Id        primitive.ObjectID `bson:"_id" json:"id"`
	ListId    primitive.ObjectID `bson:"listId" json:"listId"`
	Number    string             `bson:"number" json:"number"`
	Status    string             `bson:"status" json:"status"`
	CreatedAt time.Time          `bson:"createdAt" json:"createdAt"`
}

func NewSubscription() *Subscription {
	return &Subscription{
		Id:        primitive.NewObjectID(),
		Status:    PENDING,
		CreatedAt: time.Now(),
	}
}
