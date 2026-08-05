package product

import (
	"go/order-api/configs"
	"go/order-api/pkg/request"
	"go/order-api/pkg/response"
	"net/http"
	"strconv"
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
	router.HandleFunc("GET /product/{id}", handler.GetByID())
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

func (handler *ProductHandler) GetByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := r.PathValue("id")
		id, err := strconv.ParseUint(idStr, 10, 64)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		product, err := handler.ProductRepository.FindByID(uint(id))
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		response.Json(w, product, http.StatusCreated)
	}
}
