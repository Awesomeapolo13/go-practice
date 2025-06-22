package api

import (
	"bin/bins"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

const allowedStatusCode = 200

const GET = "GET"
const POST = "POST"
const PUT = "PUT"
const DELETE = "DELETE"

const apiRootUrl = "https://api.jsonbin.io/v3"

type ConfigInterface interface {
	GetAccessKey() string
	GetMasterKey() string
}

type StorageInterface interface {
	AddBin(bin bins.Bin)
	FindAllBins() *bins.BinList
	ToByteSlice() ([]byte, error)
}

type API struct {
	config  ConfigInterface
	storage StorageInterface
	client  *http.Client
}

func NewAPI(config ConfigInterface, storage StorageInterface) *API {
	return &API{
		config:  config,
		storage: storage,
		client:  &http.Client{},
	}
}

func (api *API) GetBin(id string) {
	// Заглушка
}

func (api *API) CreateBin(binName string, data []byte) (bins.Bin, error) {
	url := apiRootUrl + "/b"
	req, err := http.NewRequest(POST, url, bytes.NewBuffer(data))
	if err != nil {
		panic("Failed to create request: " + err.Error())
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-Master-Key", api.config.GetMasterKey())
	req.Header.Add("X-Access-Key", api.config.GetAccessKey())
	req.Header.Add("X-Bin-Name", binName)

	resp, err := api.client.Do(req)
	if err != nil {
		fmt.Println("Failed request " + url + ": " + err.Error())
		return bins.Bin{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != allowedStatusCode {
		return bins.Bin{}, errors.New("Wrong response status: " + resp.Status)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return bins.Bin{}, errors.New("Failed to read response body: " + err.Error())
	}

	var bin bins.Bin
	err = json.Unmarshal(body, &bin)
	if err != nil {
		return bins.Bin{}, errors.New("Failed to unmarshal response body: " + err.Error())
	}
	bin.Name = binName
	api.storage.AddBin(bin)

	return bin, nil
}

func (api *API) UpdateBin(bin bins.Bin) {
	// Заглушка
}

func (api *API) DeleteBin(bin bins.Bin) {
	// Заглушка
}

func (api *API) LoadToCloud(bins bins.BinList) {
	// Пока что заглушка
}

func (api *API) LoadFromCloud() bins.BinList {
	// Пока что заглушка
	return bins.BinList{}
}

func (api *API) addToStorage(bins.Bin) {

}
