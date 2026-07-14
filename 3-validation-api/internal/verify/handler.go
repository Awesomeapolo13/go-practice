package verify

import (
	"go/validation-api/configs"
	"go/validation-api/pkg/response"
	"net/http"
)

type VerificationHandlerDeps struct {
	*configs.Config
}

type VerificationHandler struct {
	*configs.Config
}

func NewVerificationHandler(router *http.ServeMux, deps VerificationHandlerDeps) {
	handler := &VerificationHandler{
		Config: deps.Config,
	}

	router.HandleFunc("POST /send", handler.Send())
	router.HandleFunc("GET /verify/{hash}", handler.Verify())
}

func (h *VerificationHandler) Send() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res := SendVerificationResponse{
			Success: true,
		}
		// Тут будет отправка письма
		response.Json(w, res, http.StatusOK)
	}
}

func (h *VerificationHandler) Verify() http.HandlerFunc {
	return func(w http.ResponseWriter, request *http.Request) {
		res := VerificationResponse{
			Success: true,
			Message: "The email verified successfully.",
		}
		// Тут будет логика подтверждения почты
		response.Json(w, res, http.StatusOK)
	}
}
