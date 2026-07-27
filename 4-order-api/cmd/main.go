package main

import (
	"fmt"
	"go/order-api/configs"
	"go/order-api/internal/auth"
	"go/order-api/internal/product"
	"go/order-api/pkg/db"
	"net/http"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	conf := configs.LoadConfig()
	_ = db.NewDb(conf)
	runAutoMigrations()

	router := http.NewServeMux()

	auth.NewAuthHandler(router, auth.AuthenticationHandlerDeps{
		Config: conf,
	})

	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	fmt.Println("Server is listening on port 8081")
	server.ListenAndServe()
}

func runAutoMigrations() {
	database, err := gorm.Open(postgres.Open(os.Getenv("DB_DSN")), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = database.AutoMigrate(&product.Product{})
	if err != nil {
		panic(err)
	}
}
