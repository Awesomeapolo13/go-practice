package file

import (
	"encoding/json"
	"fmt"
	"os"
)

func ReadFile(name string) ([]byte, error) {
	data, err := os.ReadFile(name)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return data, nil
}

func IsJSON(fileName string) bool {
	var js map[string]interface{}

	return json.Unmarshal([]byte(fileName), &js) == nil
}

func WriteFile(content []byte, name string) {
	if !IsJSON(name) {
		fmt.Println(name + " not json file")
	}

	file, err := os.Create(name)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	_, err = file.Write(content)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Wrote successfully")
}
