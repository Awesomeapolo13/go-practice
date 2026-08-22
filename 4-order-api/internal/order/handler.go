package order

import (
	"go/order-api/configs"
	"go/order-api/pkg/middleware"
	"net/http"
)

type OrderHandlerDeps struct {
	*configs.Config
	OrderRepository *OrderRepository
}

type OrderHandler struct {
	*configs.Config
	OrderRepository *OrderRepository
}

func NewOrderHandler(router *http.ServeMux, deps OrderHandlerDeps) {
	handler := &OrderHandler{
		Config:          deps.Config,
		OrderRepository: deps.OrderRepository,
	}

	router.Handle("POST /order", middleware.IsAuthed(handler.CreateOrder(), deps.Config))
	router.Handle("GET /order/{id}", middleware.IsAuthed(handler.GetOrder(), deps.Config))
	router.Handle("GET /my-orders", middleware.IsAuthed(handler.GetOrdersByUser(), deps.Config))
}

func (h *OrderHandler) CreateOrder() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (h *OrderHandler) GetOrder() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (h *OrderHandler) GetOrdersByUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
