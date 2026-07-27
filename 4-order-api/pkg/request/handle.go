package request

import (
	"go/order-api/pkg/response"
	"net/http"
)

func HandleBody[T any](w *http.ResponseWriter, r *http.Request) (*T, error) {
	payload, err := Decode[T](r.Body)
	if err != nil {
		response.Json(*w, err, http.StatusBadRequest)
		return nil, err
	}

	err = IsValid(payload)
	if err != nil {
		response.Json(*w, err.Error(), http.StatusBadRequest)
		return nil, err
	}

	return &payload, nil
}
