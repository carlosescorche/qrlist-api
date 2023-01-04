package counter

import "go.mongodb.org/mongo-driver/bson/primitive"

func GetCounter(listId primitive.ObjectID) (*Counter, error) {
	return getCounter(listId)
}
