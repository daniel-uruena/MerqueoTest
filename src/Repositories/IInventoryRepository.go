package Repositories

import "../Models"

type IInventoryRepository interface {
	GetInventory() ([]Models.Inventory, error)
	GetInventoryOfProduct(idProduct int) ([]Models.Inventory, error)
	GetInventoryOfProductAndDate(idProduct int, deliveryDate string) (Models.Inventory, error)
	GetInventoryByDate(date string) ([]Models.Inventory, error)
}
