package main

import (
	"fmt"
	"go/order-api/configs"
	"go/order-api/internal/auth"
	"go/order-api/internal/order"
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
	app := App()
	server := http.Server{
		Addr:    ":8081",
		Handler: app,
	}

	fmt.Println("Server is listening on port 8081")
	server.ListenAndServe()
}

func runAutoMigrations() {
	database, err := gorm.Open(postgres.Open(os.Getenv("DB_DSN")), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = database.AutoMigrate(&product.Product{}, &user.User{}, &order.Order{})
	if err != nil {
		panic(err)
	}
}

func App() http.Handler {
	conf := configs.LoadConfig()
	dataBase := db.NewDb(conf)
	runAutoMigrations()

	router := http.NewServeMux()

	// Repository
	productRepo := product.NewProductRepository(dataBase)
	userRepo := user.NewUserRepository(dataBase)
	orderRepo := order.NewOrderRepository(dataBase)

	// Services
	authService := auth.NewAuthService(userRepo)
	notificatorService := notification.NewNotificator()

	// Handlers
	auth.NewAuthHandler(router, auth.AuthenticationHandlerDeps{
		Config:      conf,
		AuthService: authService,
		Notificator: notificatorService,
	})
	product.NewProductHandler(router, product.ProductHandlerDeps{
		Config:            conf,
		ProductRepository: productRepo,
	})
	order.NewOrderHandler(router, order.OrderHandlerDeps{
		Config:            conf,
		OrderRepository:   orderRepo,
		UserRepository:    userRepo,
		ProductRepository: productRepo,
	})

	stack := middleware.Chain(
		middleware.CORS,
		middleware.Logging,
	)

	return stack(router)
}
