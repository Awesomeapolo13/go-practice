package api

import (
	"bin/bins"
	"bin/config"
)

type API struct {
	config config.Config
}

func NewAPI(config config.Config) *API {
	return &API{
		config: config,
	}
}

func (api *API) LoadToCloud(bins bins.BinList) {
	// Пока что заглушка
}

func (api *API) LoadFromCloud() bins.BinList {
	// Пока что заглушка
	return bins.BinList{}
}
