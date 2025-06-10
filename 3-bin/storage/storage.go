package storage

import (
	"3_bin/bins"
	"3_bin/file"
	"encoding/json"
	"fmt"
	"time"
)

const storageFileName = "storage.json"

type Storage struct {
	Bins      *bins.BinList `json:"bins"`
	UpdatedAt time.Time
}

func NewStorage() *Storage {
	if !file.IsJSON(storageFileName) {
		panic("Invalid extension of storage file " + storageFileName + ". Only JSON files are supported.")
	}
	fileContent, err := file.ReadFile(storageFileName)
	if err != nil {
		return &Storage{
			Bins:      &bins.BinList{},
			UpdatedAt: time.Now(),
		}
	}
	var storage Storage
	err = json.Unmarshal(fileContent, &storage)
	if err != nil {
		fmt.Println("Could not unmarshal storage file")

		return &Storage{
			Bins:      &bins.BinList{},
			UpdatedAt: time.Now(),
		}
	}

	return &storage
}

func (storage *Storage) AddBin(bin bins.Bin) {
	binsList := storage.Bins.Bins
	binsList = append(binsList, bin)
	storage.Bins.Bins = binsList
	storage.UpdatedAt = time.Now()
	data, err := storage.ToByteSlice()
	if err != nil {
		fmt.Println("Could not marshal " + storageFileName)
	}
	err = file.WriteFile(data, storageFileName)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func (storage *Storage) FindAllBins() *bins.BinList {
	return storage.Bins
}

func (storage *Storage) ToByteSlice() ([]byte, error) {
	fileContent, err := json.Marshal(storage)
	if err != nil {
		return nil, err
	}

	return fileContent, nil
}
