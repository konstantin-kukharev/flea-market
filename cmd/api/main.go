package main

import (
	"flea-market/cmd/api/handler"
	"log"
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

	err := s.ListenAndServe()
	log.Fatal(err.Error())
}
