package Repositories

import "../Models"

type IInventoryRepository interface {
	GetInventory() ([]Models.Inventory, error)
	GetInventoryOfProduct(idProduct int) ([]Models.Inventory, error)
}