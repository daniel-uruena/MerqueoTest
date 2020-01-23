package Models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Provider struct {
	Id primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	IdProvider int `json:"id,omitempty" bson:"id,omitempty"`
	Name string `json:"name,omitempty" bson:"name,omitempty"`
	Products []Product `json:"products,omitempty" bson:"products,omitempty"`
}
