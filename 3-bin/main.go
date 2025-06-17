package main

import (
	binApi "bin/api"
	appConfig "bin/config"
	"bin/file"
	"bin/storage"
	"fmt"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Could not load .env file")
	}
	appConfig := appConfig.NewConfig()
	fileSrv := file.NewFile()
	storageDb := storage.NewStorage(fileSrv)
	api := binApi.NewAPI(appConfig, storageDb)
	fmt.Println(api)
}
