package main

import (
	"../Repositories"
	"fmt"
	"testing"
)

func setupProviderRepository() Repositories.ProviderRepository {
	providerRepository := Repositories.ProviderRepository{
		Server:          "mongodb://localhost:27017",
		Database:        "Merqueo",
		Collection:      "provider",
	}
	return providerRepository
}

func TestGetProviders(t *testing.T)  {
	providerRepository := setupProviderRepository()
	providers, err := providerRepository.GetProviders()
	if err != nil {
		t.Error(err.Error())
	}
	if providers != nil {
		fmt.Println("Success")
	}
}

func TestGetProviderById(t *testing.T) {
	providerRepository := setupProviderRepository()
	provider, err := providerRepository.GetProviderById(2)
	if err != nil {
		t.Error(err.Error())
	}
	if provider.IdProvider > 0 {
		fmt.Println("Success")
	}
}

func TestGetProvidersByProduct(t *testing.T)  {
	providerRepository := setupProviderRepository()
	providers, err := providerRepository.GetProvidersByProduct(2)
	if err != nil {
		t.Error(err.Error())
	}
	if providers != nil {
		fmt.Println("Success")
	}
}