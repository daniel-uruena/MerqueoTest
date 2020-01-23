package Models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Order struct {
	Id primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	IdOrder int `json:"id,omitempty" bson:"id,omitempty"`
	Priority int `json:"priority,omitempty" bson:"priority,omitempty"`
	Address string `json:"address,omitempty" bson:"address,omitempty"`
	User string `json:"user,omitempty" bson:"user,omitempty"`
	Products []Product `json:"products,omitempty" bson:"products,omitempty"`
	DeliveryDate string `json:"deliveryDate,omitempty" bson:"deliveryDate,omitempty"`
}