package main

import (
	binApi "bin/api"
	"bin/bins"
	appConfig "bin/config"
	"bin/file"
	"bin/storage"
	"fmt"
	"github.com/joho/godotenv"
)

type StorageInterface interface {
	AddBin(bin bins.Bin)
	FindAllBins() *bins.BinList
	ToByteSlice() ([]byte, error)
}

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Could not load .env file")
	}
	appConfig := appConfig.NewConfig()
	api := binApi.NewAPI(*appConfig)
	fileSrv := file.NewFile()
	storage.NewStorage(fileSrv)
	fmt.Println(api)
}
