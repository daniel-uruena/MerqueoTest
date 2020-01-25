package Models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Provider struct {
	Id         primitive.ObjectID `json:"-" bson:"_id,omitempty"`
	IdProvider int                `json:"idProvider,omitempty" bson:"id,omitempty"`
	Name       string             `json:"name,omitempty" bson:"name,omitempty"`
	Products   []Product          `json:"products,omitempty" bson:"products,omitempty"`
}
