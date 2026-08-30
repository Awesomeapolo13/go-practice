package main

import (
	"bytes"
	"encoding/json"
	"go/order-api/internal/order"
	"go/order-api/internal/product"
	"go/order-api/internal/user"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	ExpectedProductID          = 1
	ExpectedProductName        = "Sugar"
	ExpectedProductDescription = "Product Description"
)

const (
	UserPhone            = "79136661111"
	UserSessionId        = "1111"
	UserVerificationCode = "1111"
)

const ExpectedOrderId = 1

func initDB() *gorm.DB {
	err := godotenv.Load("cmd/.env")
	if err != nil {
		panic(err)
	}
	db, err := gorm.Open(postgres.Open(os.Getenv("DB_DSN")), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}

func initData(db *gorm.DB) {
	db.Create(&product.Product{
		Name:        ExpectedProductName,
		Description: ExpectedProductDescription,
		Images:      []string{},
	})
	db.Create(&user.User{
		PhoneNumber:      UserPhone,
		SessionId:        UserSessionId,
		VerificationCode: UserVerificationCode,
	})
}

func removeData(db *gorm.DB) {
	db.Unscoped().
		Where("phone_number = ?", UserPhone).
		Delete(&user.User{})
	db.Unscoped().
		Where("id = ?", ExpectedProductID).
		Delete(&product.Product{})
	db.Unscoped().
		Where("id = ?", ExpectedOrderId).
		Delete(&order.Order{})
}

func TestCreateOrderSuccess(t *testing.T) {
	// Prepare
	db := initDB()
	initData(db)
	ts := httptest.NewServer(App())
	defer ts.Close()

	// Получить токен и добавить его в заголовки к запросу

	tomorrow := time.Now().Add(time.Hour * 24)
	data, err := json.Marshal(&order.CreateOrderRequest{
		OrderDate:  tomorrow,
		IsDelivery: true,
		IsExpress:  true,
		Products:   []uint{ExpectedProductID},
	})
	if err != nil {
		t.Fatalf("Error while preparing test data: %v", err)
	}

	resp, err := http.Post(ts.URL+"/order", "application/json", bytes.NewReader(data))
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("login failed, expected %d, got %d", http.StatusOK, resp.StatusCode)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	var resData order.Order
	err = json.Unmarshal(body, &resData)
	if err != nil {
		t.Fatal(err)
	}

	removeData(db)
}
