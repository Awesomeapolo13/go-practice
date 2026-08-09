package middleware

import (
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.SetFormatter(&log.JSONFormatter{})

		start := time.Now()
		wrapper := &WrapperWriter{
			ResponseWriter: w,
			StatusCode:     http.StatusOK,
		}

		next.ServeHTTP(wrapper, r)
		log.WithFields(log.Fields{
			"method":         r.Method,
			"path":           r.URL.Path,
			"status_code":    wrapper.StatusCode,
			"execution_time": time.Since(start),
		}).Info("Handled request")
	})
}
