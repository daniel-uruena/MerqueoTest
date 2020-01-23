package Repositories

import (
	"../Models"
)

type IProviderRepository interface {
	GetProviders() ([]Models.Provider, error)
	GetProviderById(idProvider int) (Models.Provider, error)
	GetProviderByProduct(idProduct int) ([]Models.Provider, error)
}