package bins

import (
	"errors"
	"time"
)

type Bin struct {
	Id        string
	Private   bool
	CreatedAt time.Time
	Name      string
}

type BinList struct {
	Bins []Bin
}

func NewBin(id, name string, private bool) (*Bin, error) {
	if id == "" {
		return nil, errors.New("Empty ID")
	}
	if name == "" {
		return nil, errors.New("Empty name")
	}

	newBin := Bin{
		Id:        id,
		Private:   private,
		CreatedAt: time.Now(),
		Name:      name,
	}

	return &newBin, nil
}

func NewBinList(bins []Bin) *BinList {
	return &BinList{Bins: bins}
}
