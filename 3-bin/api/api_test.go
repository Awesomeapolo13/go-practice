package api_test

import (
	binApi "bin/api"
	appConfig "bin/config"
	"bin/file"
	"bin/storage"
	"testing"
)

// Покрыть тестами:
// - создание бин (создать и удалить)
// - обновление бин (создать, обновить, удалить)
// - получение созданного бин (создать, получить, удалить)
// - удаление бин (создать и удалить)
//

func TestCreateBin(t *testing.T) {
	// Arrange
	expectedName := "Created Bin"
	binFile := "./create_bin.json"
	data := extractAndCheckData(binFile)
	appConfig := appConfig.NewConfig()
	storageDb := storage.NewStorage(file.NewFile())
	api := binApi.NewAPI(appConfig, storageDb)
	// Act
	bin, err := api.CreateBin(expectedName, data)
	// Assert
	if err != nil {
		t.Errorf("CreateBin failed. Got an error %v", err)
	}
	if bin.Name != expectedName {
		t.Errorf("Wrong Bin name. Got an error %s, but %s was expected", bin.Name, expectedName)
	}
	// Tear down
	err = api.DeleteBin(bin.Id)
	if err != nil {
		panic(err)
	}
}

func TestUpdateBin(t *testing.T) {
	// Arrange
	expectedName := "Updated Bin"
	binFileToCreate := "./create_bin.json"
	binFileToUpdate := "./update_bin.json"
	createData := extractAndCheckData(binFileToCreate)
	updateData := extractAndCheckData(binFileToUpdate)
	appConfig := appConfig.NewConfig()
	storageDb := storage.NewStorage(file.NewFile())
	api := binApi.NewAPI(appConfig, storageDb)
	bin, err := api.CreateBin(expectedName, createData)
	if err != nil {
		panic(err)
	}
	// Act
	err = api.UpdateBin(bin.Id, updateData)
	// Assert
	if err != nil {
		t.Errorf("UpdateBin failed. Got an error %v", err)
	}
	if bin.Name != expectedName {
		t.Errorf("Wrong Bin name. Got an error %s, but %s was expected", bin.Name, expectedName)
	}
	// Tear down
	err = api.DeleteBin(bin.Id)
	if err != nil {
		panic(err)
	}
}

func extractAndCheckData(fileName string) []byte {
	fileSrv := file.NewFile()
	if !fileSrv.IsJSON(fileName) {
		panic("Wrong file format. JSON format is only permitted")
	}

	data, err := fileSrv.ReadFile(fileName)
	if err != nil {
		panic("Could not read a file with data. Check the file and try again")
	}

	return data
}
