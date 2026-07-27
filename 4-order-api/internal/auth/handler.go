package auth

import (
	"fmt"
	"go/order-api/configs"
	"go/order-api/pkg/request"
	"go/order-api/pkg/response"
	"net/http"
)

type AuthenticationHandlerDeps struct {
	*configs.Config
}

type AuthenticationHandler struct {
	*configs.Config
}

func NewAuthHandler(router *http.ServeMux, deps AuthenticationHandlerDeps) {
	handler := &AuthenticationHandler{
		Config: deps.Config,
	}

	router.HandleFunc("POST /auth/login", handler.AuthByPhone())
	router.HandleFunc("GET /auth/verify-code", handler.Verify())
}

func (h *AuthenticationHandler) AuthByPhone() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := request.HandleBody[AuthRequest](&w, r)
		if err != nil {
			return
		}
		fmt.Println(body)

		// Тут будет авторизация
		response.Json(w, r, http.StatusOK)
	}
}

func (h *AuthenticationHandler) Verify() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := request.HandleBody[VerifyRequest](&w, r)
		if err != nil {
			return
		}
		fmt.Println(body)
		// Здесь будет верификация кода
		response.Json(w, r, http.StatusOK)
	}
}
