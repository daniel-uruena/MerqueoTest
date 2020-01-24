package Controllers

import (
	"../Repositories"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type InventoryController struct {
	BaseController
	InventoryRepository Repositories.InventoryRepository
}

func (this *InventoryController) GetInventory(response http.ResponseWriter, request *http.Request) {
	this.setResponseHeaders(response)
	inventory, err := this.InventoryRepository.GetInventory()
	if err != nil {
		this.returnServerError(err, response)
		return
	}
	if inventory == nil {
		this.returnNotFound(response)
		return
	}
	_ = json.NewEncoder(response).Encode(inventory)
}

func (this *InventoryController) GetInventoryByProduct(response http.ResponseWriter, request *http.Request) {
	this.setResponseHeaders(response)
	params := mux.Vars(request)
	id, _ := strconv.Atoi(params["idProduct"])
	productInventory, err := this.InventoryRepository.GetInventoryOfProduct(id)
	if err != nil {
		this.returnServerError(err, response)
		return
	}
	if productInventory == nil {
		this.returnNotFound(response)
		return
	}
	_ = json.NewEncoder(response).Encode(productInventory)
}