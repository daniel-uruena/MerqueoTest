package Repositories

import "../Models"

type IOrderRepository interface {
	GetOrders() ([]Models.Order, error)
	GetOrderById(idOrder int) (Models.Order, error)
	GetOrdersByProduct(idProduct int) ([]Models.Order, error)
	GetOrderByDate(date string) ([]Models.Order, error)
}
