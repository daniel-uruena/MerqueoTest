package Repositories

import (
	"../Models"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

type ProviderRepository struct {
	MongoRepository
	Server string
	Database string
	Collection string
}

func (this ProviderRepository) GetProviders() ([]Models.Provider, error) {
	return this.GetProvidersByProduct(0)
}

func (this ProviderRepository) GetProviderById(idProvider int) (Models.Provider, error) {
	this.getClient(this.Server)
	collection := this.client.Database(this.Database).Collection(this.Collection)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	var provider Models.Provider
	err := collection.FindOne(ctx, Models.Provider{IdProvider:idProvider}).Decode(&provider)
	if err != nil {
		return Models.Provider{}, err
	}
	return provider, nil
}

func (this ProviderRepository) GetProvidersByProduct(idProduct int) ([]Models.Provider, error) {
	this.getClient(this.Server)
	collection := this.client.Database(this.Database).Collection(this.Collection)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cursor, err := collection.Find(ctx, this.validateFilter(idProduct))
	if err != nil {
		return nil, err
	}
	var providers []Models.Provider
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var provider Models.Provider
		_ = cursor.Decode(&provider)
		providers = append(providers, provider)
	}
	return providers, nil
}

func (this ProviderRepository) validateFilter(idProduct int) interface{} {
	if idProduct == 0 {
		return bson.M{}
	}
	filter := bson.D{
		{"products", bson.D{
			{ "$elemMatch", bson.D{
				{"id", idProduct },
			}},
		} },
	}
	return filter
}