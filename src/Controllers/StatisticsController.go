package Controllers

import (
	"../Models"
	"../Repositories"
	"encoding/json"
	"github.com/ahmetb/go-linq"
	"github.com/gorilla/mux"
	"net/http"
	"regexp"
)

type StatisticsController struct {
	BaseController
	OrderRepository Repositories.OrderRepository
}

func (this *StatisticsController) GetBestSoldByDate(response http.ResponseWriter, request *http.Request) {
	this.setResponseHeaders(response)
	params := mux.Vars(request)
	date, _ := params["deliveryDate"]
	matched, err := regexp.MatchString(`^\d{4}-\d{2}-\d{2}$`, date)
	if err != nil {
		this.returnBadRequest(err.Error(), response)
		return
	} else if !matched {
		this.returnBadRequest("The format of date must be YYYY-MM-DD (Ex. 2020-01-27)", response)
		return
	}
	orders, err := this.OrderRepository.GetOrderByDate(date)
	if err != nil {
		this.returnServerError(err, response)
		return
	}
	if orders == nil {
		this.returnNotFound(response)
		return
	}
	products := orderProductsBySelling(orders, false)
	_ = json.NewEncoder(response).Encode(products)
}

func (this *StatisticsController) GetLessSoldByDate(response http.ResponseWriter, request *http.Request) {
	this.setResponseHeaders(response)
	params := mux.Vars(request)
	date, _ := params["deliveryDate"]
	matched, err := regexp.MatchString(`^\d{4}-\d{2}-\d{2}$`, date)
	if err != nil {
		this.returnBadRequest(err.Error(), response)
		return
	} else if !matched {
		this.returnBadRequest("The format of date must be YYYY-MM-DD (Ex. 2020-01-27)", response)
		return
	}
	orders, err := this.OrderRepository.GetOrderByDate(date)
	if err != nil {
		this.returnServerError(err, response)
		return
	}
	if orders == nil {
		this.returnNotFound(response)
		return
	}
	products := orderProductsBySelling(orders, true)
	_ = json.NewEncoder(response).Encode(products)
}

func orderProductsBySelling(orders []Models.Order, ascending bool) []Models.Product {
	var orderedProducts []Models.Product
	for _, order := range orders {
		linq.From(order.Products).ForEachT(func(product Models.Product) {
			var existingProduct []Models.Product
			linq.From(orderedProducts).WhereT(func(orderedProduct Models.Product) bool {
				return orderedProduct.Id == product.Id
			}).ToSlice(&existingProduct)
			if len(existingProduct) > 0 {
				for _, p := range existingProduct {
					p.Quantity += product.Quantity
				}
			} else {
				orderedProducts = append(orderedProducts, product)
			}
		})
	}
	linq.From(orderedProducts).SortT(func(productA, productB Models.Product) bool {
		if ascending {
			return productA.Quantity < productB.Quantity
		} else {
			return productA.Quantity > productB.Quantity
		}
	}).ToSlice(&orderedProducts)
	return orderedProducts
}
