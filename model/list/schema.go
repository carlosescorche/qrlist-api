package list

import (
	"fmt"
	"time"

	"github.com/carlosescorche/qrlist/utils/validator"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type List struct {
	Id          primitive.ObjectID `json:"id" bson:"_id"`
	Name        string             `bson:"name" json:"name"`
	Description string             `bson:"description" json:"description"`
	CreatedAt   time.Time          `json:"createdAt" bson:"createdAt" validate:"required"`
	UpdatedAt   time.Time          `json:"updatedAt" bson:"updatedAt" validate:"required"`
}

func NewList() *List {
	return &List{
		Id:        primitive.NewObjectID(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (l *List) ValidateSchema() error {
	if errs, ok := validator.ValidateStruct(l); !ok {
		return fmt.Errorf("%v", errs)
	}
	return nil
}
