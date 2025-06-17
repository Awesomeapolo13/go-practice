package main

import (
	"bin/bins"
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
}
