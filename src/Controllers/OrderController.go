package Controllers

import (
	"../Repositories"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type OrderController struct {
	BaseController
	OrderRepository Repositories.IOrderRepository
}

func (this *OrderController) GetOrders(response http.ResponseWriter, request *http.Request) {
	this.setResponseHeaders(response)
	orders, err := this.OrderRepository.GetOrders()
	if err != nil {
		this.returnServerError(err, response)
		return
	}
	if orders == nil {
		this.returnNotFound(response)
		return
	}
	_ = json.NewEncoder(response).Encode(orders)
}

func (this *OrderController) GetOrderById(response http.ResponseWriter, request *http.Request) {
	this.setResponseHeaders(response)
	params := mux.Vars(request)
	id, _ := strconv.Atoi(params["idOrder"])
	order, err := this.OrderRepository.GetOrderById(id)
	if err != nil {
		this.returnServerError(err, response)
		return
	}
	if order.IdOrder == 0 {
		this.returnNotFound(response)
		return
	}
	_ = json.NewEncoder(response).Encode(order)
}

func (this *OrderController) GetOrdersByProduct(response http.ResponseWriter, request *http.Request) {
	this.setResponseHeaders(response)
	params := mux.Vars(request)
	id, _ := strconv.Atoi(params["idProduct"])
	orders, err := this.OrderRepository.GetOrdersByProduct(id)
	if err != nil {
		this.returnServerError(err, response)
		return
	}
	if orders == nil {
		this.returnNotFound(response)
		return
	}
	_ = json.NewEncoder(response).Encode(orders)
}