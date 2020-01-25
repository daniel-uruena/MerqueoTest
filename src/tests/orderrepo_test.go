package main

import (
	"../Repositories"
	"fmt"
	"testing"
)

func setupOrderRepository() Repositories.OrderRepository {
	orderRepository := Repositories.OrderRepository{
		Server:          "mongodb://localhost:27017",
		Database:        "Merqueo",
		Collection:      "order",
	}
	return orderRepository
}

func TestGetOrders (t *testing.T) {
	orderRepository := setupOrderRepository()
	orders, err := orderRepository.GetOrders()
	if err != nil {
		t.Error(err.Error())
	}
	if orders != nil {
		fmt.Println("Success")
	}
}

func TestGetOdersByProduct (t *testing.T) {
	orderRepository := setupOrderRepository()
	orders, err := orderRepository.GetOrdersByProduct(1)
	if err != nil {
		t.Error(err.Error())
	}
	if orders != nil {
		fmt.Println("Success")
	} else {
		t.Error("orders was not found")
	}
}

func TestGetOrderById(t *testing.T) {
	orderRepository := setupOrderRepository()
	order, err := orderRepository.GetOrderById(3)
	if err != nil {
		t.Error(err.Error())
	}
	if order.IdOrder > 0 {
		fmt.Println("Success")
	}
}
