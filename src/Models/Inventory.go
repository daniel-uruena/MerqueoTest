package Models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Inventory struct {
	Id primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	IdProduct int `json:"id,omitempty" bson:"id,omitempty"`
	Date time.Time `json:"date,omitempty" bson:"date,omitempty"`
	Quantity int `json:"quantity,omitempty" bson:"quantity,omitempty"`
}