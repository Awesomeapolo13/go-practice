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
	FindBinById(id string) (*bins.Bin, error)
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

func (api *API) CreateBin(binName string, data []byte) (*bins.Bin, error) {
	url := apiRootUrl + "/b"
	req, err := api.prepareRequest(url, POST, data)
	if err != nil {
		return nil, err
	}
	req.Header.Add("X-Bin-Name", binName)

	body, err := api.getResponse(url, req)
	if err != nil {
		return nil, err
	}

	var bin bins.Bin
	err = json.Unmarshal(body, &bin)
	if err != nil {
		return nil, errors.New("Failed to unmarshal response body: " + err.Error())
	}
	bin.Name = binName
	api.storage.AddBin(bin)

	return &bin, nil
}

func (api *API) UpdateBin(binId string, data []byte) error {
	url := apiRootUrl + "/b/" + binId
	req, err := api.prepareRequest(url, PUT, data)
	if err != nil {
		return err
	}
	_, err = api.getResponse(url, req)
	if err != nil {
		return err
	}

	return nil
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

func (api *API) prepareRequest(url, method string, data []byte) (*http.Request, error) {
	req, err := http.NewRequest(method, url, bytes.NewBuffer(data))
	if err != nil {
		return nil, errors.New("Failed to create request: " + err.Error())
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-Master-Key", api.config.GetMasterKey())
	req.Header.Add("X-Access-Key", api.config.GetAccessKey())

	return req, nil
}

func (api *API) getResponse(url string, req *http.Request) ([]byte, error) {
	resp, err := api.client.Do(req)
	if err != nil {
		fmt.Println("Failed request " + url + ": " + err.Error())
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != allowedStatusCode {
		return nil, errors.New("Wrong response status: " + resp.Status)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("Failed to read response body: " + err.Error())
	}

	return body, nil
}
