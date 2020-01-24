package main

import (
	"./Controllers"
	"./Models"
	"./Repositories"
)

type ApiContext struct {
	InventoryController Controllers.InventoryController
	OrderController Controllers.OrderController
	ProviderController Controllers.ProviderController
}

func GetContext(config Models.Configuration) ApiContext {
	context := ApiContext{
		InventoryController: Controllers.InventoryController{
			InventoryRepository: Repositories.InventoryRepository{
				Server:     config.MongoServer.URI,
				Database:   config.MongoServer.DataBase,
				Collection: config.MongoServer.InventoryCollection,
			},
		},
		OrderController:Controllers.OrderController{
			OrderRepository: Repositories.OrderRepository{
				Server:     config.MongoServer.URI,
				Database:   config.MongoServer.DataBase,
				Collection: config.MongoServer.OrderCollection,
			},
		},
		ProviderController:Controllers.ProviderController{
			ProviderRepository: Repositories.ProviderRepository{
				Server:     config.MongoServer.URI,
				Database:   config.MongoServer.DataBase,
				Collection: config.MongoServer.ProviderCollection,
			},
		},
	}
	return context
}
