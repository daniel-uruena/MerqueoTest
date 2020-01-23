package main

import (
	"../Repositories"
	"fmt"
	"testing"
)

func setupInventoryRepository() Repositories.InventoryRepository {
	inventoryRepo := Repositories.InventoryRepository{
		Server: "mongodb://localhost:27017",
		Database: "Merqueo",
		Collection: "inventory",
	}
	return inventoryRepo
}

func TestGetInventory( t *testing.T) {
	inventoryRepo := setupInventoryRepository()
	inventory, err := inventoryRepo.GetInventory()
	if err != nil {
		t.Errorf(err.Error())
	}
	if inventory != nil {
		fmt.Println("Success")
	}
}

func TestGetProductInventory( t *testing.T) {
	inventoryRepo := setupInventoryRepository()
	inventory, err := inventoryRepo.GetInventoryOfProduct(5)
	if err != nil {
		t.Errorf(err.Error())
	}
	if inventory != nil {
		fmt.Println("Success")
	}
}