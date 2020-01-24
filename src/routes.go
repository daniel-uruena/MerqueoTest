package main

import (
	"github.com/gorilla/mux"
)


func RegisterRoutes(context ApiContext, prefix string) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc(prefix + "/inventory", context.InventoryController.GetInventory).Methods("GET")
	router.HandleFunc(prefix + "/inventory/{idProduct}", context.InventoryController.GetInventoryByProduct).Methods("GET")
	router.HandleFunc(prefix + "/order", context.OrderController.GetOrders).Methods("GET")
	router.HandleFunc(prefix + "/order/{idOrder}", context.OrderController.GetOrderById).Methods("GET")
	router.HandleFunc(prefix + "/order/product/{idProduct}", context.OrderController.GetOrdersByProduct).Methods("GET")
	router.HandleFunc(prefix + "/provider", context.ProviderController.GetProviders).Methods("GET")
	router.HandleFunc(prefix + "/provider/{idProvider}", context.ProviderController.GetProviderById).Methods("GET")
	router.HandleFunc(prefix + "/provider/product/{idProduct}", context.ProviderController.GetProvidersByProduct).Methods("GET")
	return router
}