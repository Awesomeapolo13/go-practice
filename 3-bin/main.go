package main

import (
	"bin/bins"
)

type StorageInterface interface {
	AddBin(bin bins.Bin)
	FindAllBins() *bins.BinList
	ToByteSlice() ([]byte, error)
}

func main() {

}
