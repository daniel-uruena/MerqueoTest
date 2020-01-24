package Controllers

import (
	"../Repositories"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type ProviderController struct {
	BaseController
	ProviderRepository Repositories.ProviderRepository
}

func (this *ProviderController) GetProviders(response http.ResponseWriter, request *http.Request) {
	this.setResponseHeaders(response)
	providers, err := this.ProviderRepository.GetProviders()
	if err != nil {
		this.returnServerError(err, response)
	}
	if providers == nil {
		this.returnNotFound(response)
	}
	_ = json.NewEncoder(response).Encode(providers)
}

func (this *ProviderController) GetProviderById(response http.ResponseWriter, request *http.Request) {
	this.setResponseHeaders(response)
	params := mux.Vars(request)
	id, _ := strconv.Atoi(params["idProvider"])
	provider, err := this.ProviderRepository.GetProviderById(id)
	if err != nil {
		this.returnServerError(err, response)
		return
	}
	if provider.IdProvider == 0 {
		this.returnNotFound(response)
		return
	}
	_ = json.NewEncoder(response).Encode(provider)
}

func (this *ProviderController) GetProvidersByProduct(response http.ResponseWriter, request *http.Request) {
	this.setResponseHeaders(response)
	params := mux.Vars(request)
	id, _ := strconv.Atoi(params["idProduct"])
	providers, err := this.ProviderRepository.GetProvidersByProduct(id)
	if err != nil {
		this.returnServerError(err, response)
		return
	}
	if providers == nil {
		this.returnNotFound(response)
		return
	}
	_ = json.NewEncoder(response).Encode(providers)
}
