package product

import (
	"go/order-api/configs"
	"go/order-api/pkg/request"
	"go/order-api/pkg/response"
	"net/http"
)

type ProductHandlerDeps struct {
	*configs.Config
	ProductRepository *ProductRepository
}

type ProductHandler struct {
	*configs.Config
	ProductRepository *ProductRepository
}

func NewProductHandler(router *http.ServeMux, deps ProductHandlerDeps) {
	handler := &ProductHandler{
		Config:            deps.Config,
		ProductRepository: deps.ProductRepository,
	}

	router.HandleFunc("POST /product", handler.Create())
}

func (handler *ProductHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := request.HandleBody[CreateProductRequest](&w, r)
		if err != nil {
			return
		}
		product := NewProduct(body.Name, body.Description, body.Images)
		createdProduct, err := handler.ProductRepository.Create(product)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		response.Json(w, createdProduct, http.StatusCreated)
	}
}
