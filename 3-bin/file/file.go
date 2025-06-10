package file

import (
	"fmt"
	"os"
	"strings"
)

const allowedExtension = ".json"

func ReadFile(name string) ([]byte, error) {
	data, err := os.ReadFile(name)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return data, nil
}

func IsJSON(fileName string) bool {
	return strings.Contains(fileName, allowedExtension)
}

func WriteFile(content []byte, name string) error {
	if !IsJSON(name) {
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
