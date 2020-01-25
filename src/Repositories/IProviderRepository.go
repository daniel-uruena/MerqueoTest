package Repositories

import (
	"../Models"
)

type IProviderRepository interface {
	GetProviders() ([]Models.Provider, error)
	GetProviderById(idProvider int) (Models.Provider, error)
	GetProvidersByProduct(idProduct int) ([]Models.Provider, error)
}