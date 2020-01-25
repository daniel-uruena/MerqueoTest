package main

import (
	"github.com/gorilla/mux"
)

func RegisterRoutes(context ApiContext, prefix string) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc(prefix+"/inventory", context.InventoryController.GetInventory).Methods("GET")
	router.HandleFunc(prefix+"/inventory/{idProduct}", context.InventoryController.GetInventoryByProduct).Methods("GET")
	router.HandleFunc(prefix+"/order", context.OrderController.GetOrders).Methods("GET")
	router.HandleFunc(prefix+"/order/{idOrder}", context.OrderController.GetOrderById).Methods("GET")
	router.HandleFunc(prefix+"/order/product/{idProduct}", context.OrderController.GetOrdersByProduct).Methods("GET")
	router.HandleFunc(prefix+"/provider", context.ProviderController.GetProviders).Methods("GET")
	router.HandleFunc(prefix+"/provider/{idProvider}", context.ProviderController.GetProviderById).Methods("GET")
	router.HandleFunc(prefix+"/provider/product/{idProduct}", context.ProviderController.GetProvidersByProduct).Methods("GET")
	router.HandleFunc(prefix+"/transporter", context.TransporterController.OrdersByTransporter).Methods("GET")
	router.HandleFunc(prefix+"/statistics/bestsold/{deliveryDate}", context.StatisticsController.GetBestSoldByDate).Methods("GET")
	router.HandleFunc(prefix+"/statistics/lesssold/{deliveryDate}", context.StatisticsController.GetLessSoldByDate).Methods("GET")
	router.HandleFunc(prefix+"/checkinventory/{idOrder}", context.CheckInventoryController.CheckInventoryToOrder).Methods("GET")
	router.HandleFunc(prefix+"/calculateinventory/{date}", context.CheckInventoryController.CalculateInventory).Methods("GET")

	return router
}
