package files

import (
	"encoding/json"
	"os"
)

type JsonDB struct {
	fileName string
}

func NewJsonDB(name string) *JsonDB {
	return &JsonDB{
		fileName: name,
	}
}

func (db *JsonDB) Read() ([]byte, error) {
	data, err := os.ReadFile(db.fileName)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (db *JsonDB) Write(content []byte) error {
	file, err := os.Create(db.fileName)
	if err != nil {
		return err
	}
	defer func(file *os.File) error {
		err := file.Close()
		if err != nil {
			return err
		}
		return nil
	}(file)

	_, err = file.Write(content)
	if err != nil {
		return err
	}
	return nil
}

func GetCollection[T any](db *JsonDB, name string) ([]T, error) {
	raw, err := db.Read()
	if err != nil {
		return []T{}, nil
	}

	var all map[string]json.RawMessage
	if err = json.Unmarshal(raw, &all); err != nil {
		return []T{}, nil
	}

	var items []T
	if err = json.Unmarshal(all[name], &items); err != nil {
		return []T{}, nil
	}
	return items, nil
}

func SetCollection[T any](db *JsonDB, name string, items []T) error {
	raw, _ := db.Read()

	all := map[string]json.RawMessage{}
	if raw != nil {
		json.Unmarshal(raw, &all)
	}

	encoded, err := json.Marshal(items)
	if err != nil {
		return err
	}
	all[name] = encoded

	data, err := json.Marshal(all)
	if err != nil {
		return err
	}
	err = db.Write(data)
	if err != nil {
		return err
	}
	return nil
}
