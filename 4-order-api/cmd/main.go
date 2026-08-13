package main

import (
	"fmt"
	"go/order-api/configs"
	"go/order-api/internal/auth"
	"go/order-api/internal/product"
	"go/order-api/internal/user"
	"go/order-api/pkg/db"
	"go/order-api/pkg/middleware"
	"go/order-api/pkg/notification"
	"net/http"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	conf := configs.LoadConfig()
	dataBase := db.NewDb(conf)
	runAutoMigrations()

	router := http.NewServeMux()
	productRepo := product.NewProductRepository(dataBase)
	userRepo := user.NewUserRepository(dataBase)

	authService := auth.NewAuthService(userRepo)
	notificatorService := notification.NewNotificator()

	auth.NewAuthHandler(router, auth.AuthenticationHandlerDeps{
		Config:      conf,
		AuthService: authService,
		Notificator: notificatorService,
	})
	product.NewProductHandler(router, product.ProductHandlerDeps{
		Config:            conf,
		ProductRepository: productRepo,
	})

	chain := middleware.Chain(
		middleware.CORS,
		middleware.Logging,
	)

	server := http.Server{
		Addr:    ":8081",
		Handler: chain(router),
	}

	fmt.Println("Server is listening on port 8081")
	server.ListenAndServe()
}

func runAutoMigrations() {
	database, err := gorm.Open(postgres.Open(os.Getenv("DB_DSN")), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = database.AutoMigrate(&product.Product{}, &user.User{})
	if err != nil {
		panic(err)
	}
}
