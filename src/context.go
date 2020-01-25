package main

import (
	"./Controllers"
	"./Models"
	"./Repositories"
)

type ApiContext struct {
	InventoryController      Controllers.InventoryController
	OrderController          Controllers.OrderController
	ProviderController       Controllers.ProviderController
	TransporterController    Controllers.TransporterController
	StatisticsController     Controllers.StatisticsController
	CheckInventoryController Controllers.CheckInventoryController
}

func GetContext(config Models.Configuration) ApiContext {
	inventoryRepository := Repositories.InventoryRepository{
		Server:     config.MongoServer.URI,
		Database:   config.MongoServer.DataBase,
		Collection: config.MongoServer.InventoryCollection,
	}
	orderRepository := Repositories.OrderRepository{
		Server:     config.MongoServer.URI,
		Database:   config.MongoServer.DataBase,
		Collection: config.MongoServer.OrderCollection,
	}
	providerRepository := Repositories.ProviderRepository{
		Server:     config.MongoServer.URI,
		Database:   config.MongoServer.DataBase,
		Collection: config.MongoServer.ProviderCollection,
	}
	context := ApiContext{
		InventoryController: Controllers.InventoryController{
			InventoryRepository: inventoryRepository,
		},
		OrderController: Controllers.OrderController{
			OrderRepository: orderRepository,
		},
		ProviderController: Controllers.ProviderController{
			ProviderRepository: providerRepository,
		},
		TransporterController: Controllers.TransporterController{
			OrderRepository: orderRepository,
		},
		StatisticsController: Controllers.StatisticsController{
			OrderRepository: orderRepository,
		},
		CheckInventoryController: Controllers.CheckInventoryController{
			InventoryRepository: inventoryRepository,
			OrderRepository:     orderRepository,
			ProviderRepository:  providerRepository,
		},
	}
	return context
}
