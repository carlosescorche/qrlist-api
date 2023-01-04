package counter

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Counter struct {
	Id        primitive.ObjectID `bson:"_id" json:"id"`
	ListId    primitive.ObjectID `bson:"listId" json:"listId"`
	Count     int                `bson:"count" json:"count"`
	CreatedAt time.Time          `bson:"createdAt" json:"createdAt"`
	ExpireOn  time.Time          `bson:"expireOn" json:"expireOn"`
}

func NewCounter(listId primitive.ObjectID) *Counter {
	return &Counter{
		ListId:    listId,
		Count:     1,
		CreatedAt: time.Now(),
		ExpireOn:  time.Now().Add(time.Minute * 5),
	}
}
