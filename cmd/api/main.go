package main

import (
	"flea-market/cmd/api/handler"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
)

func main() {
	productsHandler := handler.NewGetProducts()
	router := chi.NewRouter()
	router.Method("GET", "/api/products", productsHandler)

	s := &http.Server{
		Handler:           router,
		Addr:              "0.0.0.0:8080",
		ReadHeaderTimeout: 1 * time.Second,
	}

	s.ListenAndServe()
}
