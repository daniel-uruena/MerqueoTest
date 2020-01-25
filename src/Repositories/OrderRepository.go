package Repositories

import (
	"../Models"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

type OrderRepository struct {
	MongoRepository
	Server     string
	Database   string
	Collection string
}

func (this *OrderRepository) GetOrders() ([]Models.Order, error) {
	this.getClient(this.Server)
	return this.GetOrdersByProduct(0)
}

func (this *OrderRepository) GetOrderById(idOrder int) (Models.Order, error) {
	this.getClient(this.Server)
	collection := this.client.Database(this.Database).Collection(this.Collection)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	var order Models.Order
	err := collection.FindOne(ctx, Models.Order{IdOrder: idOrder}).Decode(&order)
	if err != nil {
		return Models.Order{}, err
	}
	return order, nil
}

func (this *OrderRepository) GetOrdersByProduct(idProduct int) ([]Models.Order, error) {
	this.getClient(this.Server)
	collection := this.client.Database(this.Database).Collection(this.Collection)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cursor, err := collection.Find(ctx, this.validateFilter(idProduct))
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	var orders []Models.Order
	for cursor.Next(ctx) {
		var result Models.Order
		_ = cursor.Decode(&result)
		orders = append(orders, result)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return orders, nil
}

func (this *OrderRepository) GetOrderByDate(date string) ([]Models.Order, error) {
	this.getClient(this.Server)
	collection := this.client.Database(this.Database).Collection(this.Collection)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cursor, err := collection.Find(ctx, Models.Order{DeliveryDate: date})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	var orders []Models.Order
	for cursor.Next(ctx) {
		var result Models.Order
		_ = cursor.Decode(&result)
		orders = append(orders, result)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return orders, nil
}

func (this OrderRepository) validateFilter(idProduct int) interface{} {
	if idProduct == 0 {
		return bson.M{}
	}
	filter := bson.D{
		{"products", bson.D{
			{"$elemMatch", bson.D{
				{"id", idProduct},
			}},
		}},
	}
	return filter
}
