package main

import (
	binApi "bin/api"
	appConfig "bin/config"
	"bin/file"
	"bin/storage"
	"flag"
	"fmt"
	"github.com/joho/godotenv"
)

func main() {
	// Зарегаться на https://jsonbin.io/ (это облачное хранилище json)
	// Реализовать методы для работы с апи, использовать в бин апи методы create a BIN, update BIN, delete BIN, read BIN
	// Бины доступны в папке бинс облачного сервиса
	// Создание бина post, чтение get, обновление put
	// Команда должна читать флаги, для create -
	//    --file - адрес файла, который хотим отправить в jsonBin (какой то локальный файл),
	//    --name - название бина (видимо для внутреннего пользования)
	// должен локально сохранить имя и идентификатор из ответа, чтобы о них осталась инфа
	// update - --file - см.выше, --id - идешка бина
	// delete - --id
	// --get --id
	// list - получить и вывести все сохраненные идентификаторы и имена
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Could not load .env file")
	}

	create := flag.Bool("create", false, "Create bin operation")
	update := flag.Bool("update", false, "Update bin operation")
	deleteOperation := flag.Bool("delete", false, "Delete bin operation")
	get := flag.Bool("get", false, "Get bin")
	list := flag.Bool("list", false, "Get list of bins")

	binsFile := flag.String("file", "", "Bin`s file")
	name := flag.String("name", "", "Bin`s name")
	id := flag.String("id", "", "Bin`s ID")

	flag.Parse()

	switch true {
	case *create:
		createBin(*binsFile, *name)
		break
	case *update:
		updateBin(*binsFile, *id)
		break
	case *deleteOperation:
		deleteBin(*id)
		break
	case *get:
		getBin(*id)
		break
	case *list:
		getBinList()
		break
	default:
		panic("Unsupported operation. Please use create, update, delete, get or list")
	}
}

func createBin(binsFile, binsName string) {
	if binsFile == "" || binsName == "" {
		panic("Please specify bins file and name")
	}

	data := extractAndCheckData("./data/" + binsFile)

	appConfig := appConfig.NewConfig()
	storageDb := storage.NewStorage(file.NewFile())
	api := binApi.NewAPI(appConfig, storageDb)
	bin, err := api.CreateBin(binsName, data)
	if err != nil {
		panic(err)
	}
	msg := fmt.Sprintf("Bin with name %s successfully created with ID %s", bin.Name, bin.Id)
	fmt.Println(msg)
}

func updateBin(binsFile, binsId string) {
	if binsFile == "" || binsId == "" {
		panic("Please specify bins file and id")
	}
	//appConfig := appConfig.NewConfig()
	storageDb := storage.NewStorage(file.NewFile())
	//api := binApi.NewAPI(appConfig, storageDb)
	binToUpdate, err := storageDb.FindBinById(binsId)
	if err != nil {
		panic(err)
	}
	fmt.Println("Updating bin", binToUpdate.Name)
	data := extractAndCheckData("./data/" + binsFile)

	appConfig := appConfig.NewConfig()
	api := binApi.NewAPI(appConfig, storageDb)
	err = api.UpdateBin(binToUpdate.Id, data)
	if err != nil {
		panic(err)
	}
}

func deleteBin(binsId string) {
	if binsId == "" {
		panic("Please specify bins id")
	}
	// Находим такой бин в локальном хранилище, если его нет то ошибка
	// Отправляем запрос в API.
	// Удаляем бин из локального хранилища
}

func getBin(binsId string) {
	if binsId == "" {
		panic("Please specify bins id")
	}
	// Находим такой бин в локальном хранилище, если его нет то ошибка
	// Получаем бин из API.
	// Выводим результат в консоль
}

func getBinList() {
	// Берем весь список из локального хранилища.
	// Выводим в консоль
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
