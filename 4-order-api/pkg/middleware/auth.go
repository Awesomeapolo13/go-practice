package middleware

import (
	"go/order-api/pkg/jwt"
	"net/http"
	"os"
	"strings"
)

func IsAuthed(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authedHeader := r.Header.Get("Authorization")
		secret := os.Getenv("SECRET")
		jwtService := jwt.NewJWT(secret)
		token := strings.TrimPrefix(authedHeader, "Bearer ")

		if token == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		jwtToken, err := jwtService.Parse(token)
		if err != nil || !jwtToken.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
