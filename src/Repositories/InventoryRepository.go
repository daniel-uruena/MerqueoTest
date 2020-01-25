package Repositories

import (
	"../Models"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

type InventoryRepository struct {
	MongoRepository
	Server     string
	Database   string
	Collection string
}

func (this InventoryRepository) GetInventory() ([]Models.Inventory, error) {
	return this.GetInventoryOfProduct(0)
}

func (this InventoryRepository) GetInventoryOfProduct(idProduct int) ([]Models.Inventory, error) {
	this.getClient(this.Server)
	collection := this.client.Database(this.Database).Collection(this.Collection)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cursor, err := collection.Find(ctx, this.validateFilter(idProduct))
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	var inventory []Models.Inventory
	for cursor.Next(ctx) {
		var result Models.Inventory
		_ = cursor.Decode(&result)
		inventory = append(inventory, result)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return inventory, nil
}

func (this InventoryRepository) GetInventoryByDate(date string) ([]Models.Inventory, error) {
	this.getClient(this.Server)
	collection := this.client.Database(this.Database).Collection(this.Collection)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cursor, err := collection.Find(ctx, Models.Inventory{Date: date})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	var inventory []Models.Inventory
	for cursor.Next(ctx) {
		var result Models.Inventory
		_ = cursor.Decode(&result)
		inventory = append(inventory, result)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return inventory, nil
}

func (this InventoryRepository) GetInventoryOfProductAndDate(idProduct int, deliveryDate string) (Models.Inventory, error) {
	this.getClient(this.Server)
	collection := this.client.Database(this.Database).Collection(this.Collection)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	var inventory Models.Inventory
	err := collection.FindOne(ctx, Models.Inventory{IdProduct: idProduct, Date: deliveryDate}).Decode(&inventory)
	if err != nil {
		return Models.Inventory{}, err
	}
	return inventory, nil
}

func (this InventoryRepository) validateFilter(idProduct int) interface{} {
	if idProduct == 0 {
		return bson.M{}
	}
	return Models.Inventory{IdProduct: idProduct}
}
