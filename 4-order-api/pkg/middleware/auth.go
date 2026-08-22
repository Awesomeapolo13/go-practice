package middleware

import (
	"context"
	"go/order-api/configs"
	"go/order-api/pkg/jwt"
	"net/http"
	"strings"
)

type key string

const (
	ContextPhoneKey key = "ContextPhoneKey"
)

func IsAuthed(next http.Handler, config *configs.Config) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authedHeader := r.Header.Get("Authorization")
		if !strings.HasPrefix(authedHeader, "Bearer ") {
			writeUnauthed(w)
			return
		}

		token := strings.TrimPrefix(authedHeader, "Bearer ")
		isValid, data := jwt.NewJWT(config.Auth.Secret).Parse(token)
		if !isValid {
			writeUnauthed(w)
			return
		}

		ctx := context.WithValue(r.Context(), ContextPhoneKey, data.Phone)
		req := r.WithContext(ctx)

		next.ServeHTTP(w, req)
	})
}

func writeUnauthed(w http.ResponseWriter) {
	w.WriteHeader(http.StatusUnauthorized)
	w.Write([]byte(http.StatusText(http.StatusUnauthorized)))
}
