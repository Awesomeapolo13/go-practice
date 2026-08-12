package auth

import (
	"go/order-api/configs"
	"go/order-api/pkg/jwt"
	"go/order-api/pkg/request"
	"go/order-api/pkg/response"
	"net/http"
)

type AuthenticationHandlerDeps struct {
	*configs.Config
	*AuthService
}

type AuthenticationHandler struct {
	*configs.Config
	*AuthService
}

func NewAuthHandler(router *http.ServeMux, deps AuthenticationHandlerDeps) {
	handler := &AuthenticationHandler{
		Config:      deps.Config,
		AuthService: deps.AuthService,
	}

	router.HandleFunc("POST /auth/login", handler.AuthByPhone())
	router.HandleFunc("POST /auth/verify-code", handler.Verify())
}

func (h *AuthenticationHandler) AuthByPhone() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := request.HandleBody[AuthRequest](&w, r)
		if err != nil {
			return
		}

		sessionId, err := h.AuthService.CreateNewSessionId(body.PhoneNumber)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		data := AuthResponse{
			SessionId: sessionId,
		}

		response.Json(w, data, http.StatusOK)
	}
}

func (h *AuthenticationHandler) Verify() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := request.HandleBody[VerifyRequest](&w, r)
		if err != nil {
			return
		}

		phone, err := h.AuthService.VerifyByCode(body.SessionId, body.Code)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		token, err := jwt.NewJWT(h.Config.Auth.Secret).Create(phone)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		data := VerifyResponse{
			Token: token,
		}

		response.Json(w, data, http.StatusOK)
	}
}
