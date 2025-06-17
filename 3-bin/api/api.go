package api

import (
	"bin/bins"
)

type ConfigInterface interface {
	GetKey() string
}

type StorageInterface interface {
	AddBin(bin bins.Bin)
	FindAllBins() *bins.BinList
	ToByteSlice() ([]byte, error)
}

type API struct {
	config  ConfigInterface
	storage StorageInterface
}

func NewAPI(config ConfigInterface, storage StorageInterface) *API {
	return &API{
		config:  config,
		storage: storage,
	}
}

func (api *API) LoadToCloud(bins bins.BinList) {
	// Пока что заглушка
}

func (api *API) LoadFromCloud() bins.BinList {
	// Пока что заглушка
	return bins.BinList{}
}
