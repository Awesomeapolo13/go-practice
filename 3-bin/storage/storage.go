package storage

import (
	"bin/bins"
	"encoding/json"
	"fmt"
	"time"
)

const storageFileName = "storage.json"

type FileInterface interface {
	ReadFile(name string) ([]byte, error)
	IsJSON(fileName string) bool
	WriteFile(content []byte, name string) error
}

type Storage struct {
	Bins      *bins.BinList `json:"bins"`
	UpdatedAt time.Time
}

type StorageWithDI struct {
	Storage
	fileSrv FileInterface
}

func NewStorage(fileSrv FileInterface) *StorageWithDI {
	if !fileSrv.IsJSON(storageFileName) {
		panic("Invalid extension of storage file " + storageFileName + ". Only JSON files are supported.")
	}
	fileContent, err := fileSrv.ReadFile(storageFileName)
	if err != nil {
		return &StorageWithDI{
			Storage: Storage{
				Bins:      &bins.BinList{},
				UpdatedAt: time.Now(),
			},
			fileSrv: fileSrv,
		}
	}
	var storage Storage
	err = json.Unmarshal(fileContent, &storage)
	if err != nil {
		fmt.Println("Could not unmarshal storage file")

		return &StorageWithDI{
			Storage: Storage{
				Bins:      &bins.BinList{},
				UpdatedAt: time.Now(),
			},
			fileSrv: fileSrv,
		}
	}

	return &StorageWithDI{
		Storage: storage,
		fileSrv: fileSrv,
	}
}

func (storage *StorageWithDI) AddBin(bin bins.Bin) {
	binsList := storage.Bins.Bins
	binsList = append(binsList, bin)
	storage.Storage.Bins.Bins = binsList
	storage.Storage.UpdatedAt = time.Now()
	data, err := storage.ToByteSlice()
	if err != nil {
		fmt.Println("Could not marshal " + storageFileName)
	}
	err = storage.fileSrv.WriteFile(data, storageFileName)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func (storage *StorageWithDI) FindAllBins() *bins.BinList {
	return storage.Storage.Bins
}

func (storage *StorageWithDI) ToByteSlice() ([]byte, error) {
	fileContent, err := json.Marshal(storage.Storage)
	if err != nil {
		return nil, err
	}

	return fileContent, nil
}
