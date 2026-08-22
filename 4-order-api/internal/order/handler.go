package order

import (
	"go/order-api/configs"
	"go/order-api/pkg/di"
	"go/order-api/pkg/middleware"
	"go/order-api/pkg/request"
	"go/order-api/pkg/response"
	"net/http"
)

type OrderHandlerDeps struct {
	*configs.Config
	OrderRepository   *OrderRepository
	UserRepository    di.IUserRepository
	ProductRepository di.IProductRepository
}

type OrderHandler struct {
	*configs.Config
	OrderRepository   *OrderRepository
	UserRepository    di.IUserRepository
	ProductRepository di.IProductRepository
}

func NewOrderHandler(router *http.ServeMux, deps OrderHandlerDeps) {
	handler := &OrderHandler{
		Config:            deps.Config,
		OrderRepository:   deps.OrderRepository,
		UserRepository:    deps.UserRepository,
		ProductRepository: deps.ProductRepository,
	}

	router.Handle("POST /order", middleware.IsAuthed(handler.CreateOrder(), deps.Config))
	router.Handle("GET /order/{id}", middleware.IsAuthed(handler.GetOrder(), deps.Config))
	router.Handle("GET /my-orders", middleware.IsAuthed(handler.GetOrdersByUser(), deps.Config))
}

func (h *OrderHandler) CreateOrder() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := request.HandleBody[CreateOrderRequest](&w, r)
		if err != nil {
			return
		}
		phone, ok := r.Context().Value(middleware.ContextPhoneKey).(string)
		if !ok {
			writeUnauthed(w)
			return
		}

		user, err := h.UserRepository.FindByPhone(phone)
		if err != nil {
			writeUnauthed(w)
			return
		}

		products, err := h.ProductRepository.FindAllByIDs(body.Products)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		order := NewOrder(
			body.OrderDate,
			body.IsDelivery,
			body.IsExpress,
			*user,
			products,
		)

		createdOrder, err := h.OrderRepository.Create(order)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		response.Json(w, createdOrder, http.StatusCreated)
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

func writeUnauthed(w http.ResponseWriter) {
	w.WriteHeader(http.StatusUnauthorized)
	w.Write([]byte(http.StatusText(http.StatusUnauthorized)))
}
