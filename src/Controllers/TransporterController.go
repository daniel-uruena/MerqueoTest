package Controllers

import (
	"../Models"
	"../Repositories"
	"encoding/json"
	"net/http"
)

type TransporterController struct {
	BaseController
	OrderRepository Repositories.IOrderRepository
}

func (this *TransporterController) OrdersByTransporter(response http.ResponseWriter, request *http.Request) {
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
	transporters := organizeOrdersToTransporters(orders)
	_ = json.NewEncoder(response).Encode(transporters)
}

func organizeOrdersToTransporters(orders []Models.Order) []*Models.Transporter {
	var transporters []*Models.Transporter
	for _, order := range orders {
		transporter := transporterByUserAndAddresOrder(transporters, order.User, order.Address)
		if transporter != nil {
			transporter.Orders = append(transporter.Orders, order)
			addProductsToTransporter(transporter, order.Products)
		} else {
			var newTransporter = Models.Transporter{
				IdTransporter: len(transporters) + 1,
				Orders: []Models.Order{
					order,
				},
				Products: order.Products,
			}
			transporters = append(transporters, &newTransporter)
		}
	}
	return transporters
}

func transporterByUserAndAddresOrder(transporters []*Models.Transporter, user string, address string) *Models.Transporter {
	for _, transporter := range transporters {
		for _, order := range transporter.Orders {
			if order.User == user && order.Address == address {
				return transporter
			}
		}
	}
	return nil
}

func addProductsToTransporter(transporter *Models.Transporter, newProducts []Models.Product) {
	existNewProduct := false
	for _, newProduct := range newProducts {
		existNewProduct = false
		for _, product := range transporter.Products {
			if product.Id == newProduct.Id {
				product.Quantity += newProduct.Quantity
				existNewProduct = true
			}
		}
		if !existNewProduct {
			transporter.Products = append(transporter.Products, newProduct)
		}
	}
}
