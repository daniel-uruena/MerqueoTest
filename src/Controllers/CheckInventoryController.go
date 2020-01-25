package Controllers

import (
	"../Models"
	"../Repositories"
	"encoding/json"
	"github.com/ahmetb/go-linq"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"time"
)

type CheckInventoryController struct {
	BaseController
	InventoryRepository Repositories.IInventoryRepository
	OrderRepository     Repositories.IOrderRepository
	ProviderRepository  Repositories.IProviderRepository
}

func (this *CheckInventoryController) CheckInventoryToOrder(response http.ResponseWriter, request *http.Request) {
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
	checkInventory := this.verifyOrderInventory(order)
	_ = json.NewEncoder(response).Encode(checkInventory)
}

func (this *CheckInventoryController) CalculateInventory(response http.ResponseWriter, request *http.Request) {
	this.setResponseHeaders(response)
	params := mux.Vars(request)
	dateStr, _ := params["date"]
	//layout := "2019-05-25"
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		this.returnBadRequest("The format of date must be YYYY-MM-DD (Ex. 2020-01-27)", response)
		return
	}
	date = date.AddDate(0, 0, -1)
	previousDateStr := date.Format("2006-01-02")
	inventory, err := this.InventoryRepository.GetInventoryByDate(previousDateStr)
	if err != nil {
		this.returnServerError(err, response)
		return
	}
	if inventory == nil {
		this.returnNotFound(response)
		return
	}
	newInventory, err := this.calculateInventory(inventory, previousDateStr, dateStr)
	if err != nil {
		this.returnServerError(err, response)
		return
	}
	_ = json.NewEncoder(response).Encode(newInventory)
}

func (this *CheckInventoryController) verifyOrderInventory(order Models.Order) checkInventory {
	var readyToDeliver []Models.Product
	var needToBeRequested []Models.Provider
	linq.From(order.Products).ForEachT(func(orderProduct Models.Product) {
		productInventory, _ := this.InventoryRepository.GetInventoryOfProductAndDate(orderProduct.Id, order.DeliveryDate)
		if productInventory.IdProduct != 0 && productInventory.Quantity >= orderProduct.Quantity {
			readyToDeliver = append(readyToDeliver, orderProduct)
		} else {
			if productInventory.IdProduct != 0 {
				readyProduct := orderProduct
				readyProduct.Quantity = productInventory.Quantity
				readyToDeliver = append(readyToDeliver, readyProduct)
			}
			providers, _ := this.ProviderRepository.GetProvidersByProduct(orderProduct.Id)
			for _, provider := range providers {
				if !linq.From(needToBeRequested).AnyWithT(func(n Models.Provider) bool {
					return n.IdProvider == provider.IdProvider
				}) {
					productProvider := Models.Provider{
						IdProvider: provider.IdProvider,
						Name:       provider.Name,
						Products:   []Models.Product{orderProduct},
					}
					if productInventory.IdProduct != 0 {
						productProvider.Products[0].Quantity = orderProduct.Quantity - productInventory.Quantity
					}
					needToBeRequested = append(needToBeRequested, productProvider)
				} else {
					var listedProvider []Models.Provider
					linq.From(needToBeRequested).WhereT(func(n Models.Provider) bool {
						return n.IdProvider == provider.IdProvider
					}).SelectT(func(n Models.Provider) Models.Provider {
						return n
					}).ToSlice(&listedProvider)

					if productInventory.IdProduct != 0 {
						orderProduct.Quantity = orderProduct.Quantity - productInventory.Quantity
					}
					listedProvider[0].Products = append(listedProvider[0].Products, orderProduct)
				}
			}
		}
	})
	return checkInventory{
		ReadyToDeliver:    readyToDeliver,
		NeedToBeRequested: needToBeRequested,
		Order:             order,
	}
}

type checkInventory struct {
	ReadyToDeliver    []Models.Product
	NeedToBeRequested []Models.Provider
	Order             Models.Order
}

func (this *CheckInventoryController) calculateInventory(inventory []Models.Inventory, previousDate string, date string) ([]Models.Inventory, error) {
	orders, err := this.OrderRepository.GetOrderByDate(previousDate)
	if err != nil {
		return nil, err
	}
	var newInventory []*Models.Inventory
	for _, order := range orders {
		for _, orderProduct := range order.Products {
			linq.From(inventory).ForEachT(func(inventoryProduct Models.Inventory) {
				existInNewInventory := linq.From(newInventory).AnyWithT(func(ni *Models.Inventory) bool {
					return ni.IdProduct == orderProduct.Id
				})
				if orderProduct.Id == inventoryProduct.IdProduct && orderProduct.Quantity < inventoryProduct.Quantity &&
					!existInNewInventory {
					inventoryProduct.Quantity = inventoryProduct.Quantity - orderProduct.Quantity
					inventoryProduct.Date = date
					newInventory = append(newInventory, &inventoryProduct)
				} else if existInNewInventory {
					var newProductInventory []*Models.Inventory
					linq.From(newInventory).WhereT(func(ni *Models.Inventory) bool {
						return ni.IdProduct == orderProduct.Id
					}).ToSlice(&newProductInventory)

					newProductInventory[0].Quantity = newProductInventory[0].Quantity - orderProduct.Quantity
				}
			})
		}
	}
	var result []Models.Inventory
	for _, newProductInventory := range newInventory {
		if newProductInventory.Quantity > 0 {
			result = append(result, *newProductInventory)
		}
	}
	return result, nil
}
