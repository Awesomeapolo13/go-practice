package file

import (
	"fmt"
	"os"
	"strings"
)

const allowedExtension = ".json"

type File struct{}

func NewFile() *File {
	return &File{}
}

func (fileSrv *File) ReadFile(name string) ([]byte, error) {
	data, err := os.ReadFile(name)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return data, nil
}

func (fileSrv *File) IsJSON(fileName string) bool {
	return strings.HasSuffix(fileName, allowedExtension)
}

func (fileSrv *File) WriteFile(content []byte, name string) error {
	if !fileSrv.IsJSON(name) {
		fmt.Println(name + " not json file")
	}

	file, err := os.Create(name)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(content)
	if err != nil {
		return err
	}
	fmt.Println("Wrote successfully")
	return nil
}
