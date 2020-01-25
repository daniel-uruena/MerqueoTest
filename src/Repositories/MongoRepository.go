package Repositories

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type MongoRepository struct {
	client *mongo.Client
}

func (this *MongoRepository) getClient(server string) {
	if this.client != nil {
		return
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	opt := options.Client().ApplyURI(server)
	var err error
	this.client, err = mongo.Connect(ctx, opt)
	if err != nil {
		fmt.Println("Error connecting to database:")
		fmt.Println(err.Error())
	}
	//return this.client
}