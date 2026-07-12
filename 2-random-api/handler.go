package main

import (
	"math/rand/v2"
	"net/http"
	"strconv"
)

type RandomIntHandler struct{}

func NewRandomIntHandler(router *http.ServeMux) {
	handler := &RandomIntHandler{}
	router.HandleFunc("/random", handler.GetRandomInt())
}

func (h *RandomIntHandler) GetRandomInt() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		minVal := 1
		maxVal := 6
		randomInt := rand.IntN(maxVal-minVal+1) + minVal

		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		_, err := w.Write([]byte(strconv.Itoa(randomInt)))
		if err != nil {
			panic(err)
		}
	}
}
