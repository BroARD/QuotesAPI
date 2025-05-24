package main

import (
	"APIWithout/internal/handlers"
	"APIWithout/internal/service"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	service := service.NewService()
	handlers := handlers.NewHandler(service)

	r := mux.NewRouter()
	r.HandleFunc("/quotes/{id}", handlers.DeleteQuoteByID).Methods("DELETE")
	r.HandleFunc("/quotes", handlers.PostQuote).Methods("POST")
	r.HandleFunc("/quotes", handlers.GetQuotes).Methods("GET")
	r.HandleFunc("/quotes/random", handlers.GetRandomQuote).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", r))
}

