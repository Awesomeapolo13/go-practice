package main

import (
	"bytes"
	"encoding/json"
	"go/order-api/internal/auth"
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

	authToken := getToken(t, ts)

	tomorrow := time.Now().Add(time.Hour * 24)
	data, err := json.Marshal(&order.CreateOrderRequest{
		OrderDate:  tomorrow,
		IsDelivery: true,
		IsExpress:  true,
		Products:   []uint{ExpectedProductID},
	})
	if err != nil {
		t.Fatalf("Error while preparing CreateOrderRequest: %v", err)
	}

	req, err := http.NewRequest("POST", ts.URL+"/order", bytes.NewBuffer(data))
	if err != nil {
		t.Fatalf("Error creating request: %v", err)
	}
	req.Header.Set("Authorization", "Bearer "+authToken)

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Creating order is failed, expected %d, got %d", http.StatusOK, resp.StatusCode)
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

func getToken(t *testing.T, ts *httptest.Server) string {
	// Получить токен и добавить его в заголовки к запросу
	verificationData, err := json.Marshal(&auth.VerifyRequest{
		SessionId: UserSessionId,
		Code:      UserVerificationCode,
	})
	if err != nil {
		t.Fatalf("Error while preparing VerifyRequest: %v", err)
	}

	verificationResp, err := http.Post(ts.URL+"/auth/verify-code", "application/json", bytes.NewReader(verificationData))
	if err != nil {
		t.Fatal(err)
	}
	if verificationResp.StatusCode != http.StatusOK {
		t.Fatalf("Could not get auth token, expected %d, got %d", http.StatusOK, verificationResp.StatusCode)
	}
	body, err := io.ReadAll(verificationResp.Body)
	if err != nil {
		t.Fatal(err)
	}

	var verificationRes auth.VerifyResponse
	err = json.Unmarshal(body, &verificationRes)
	if err != nil {
		t.Fatal(err)
	}
	if verificationRes.Token == "" {
		t.Fatal("Empty token")
	}

	return verificationRes.Token
}
