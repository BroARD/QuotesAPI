package handlers

import (
	"APIWithout/internal/service"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)


type Handler struct {
	service service.Service
}

func NewHandler(service service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) PostQuote(w http.ResponseWriter, r *http.Request) {
	var quote service.Quote
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&quote)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	h.service.CreateQuote(quote)
	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) GetQuotes(w http.ResponseWriter, r *http.Request) {
	author := r.URL.Query().Get("author")
	quotes, err := h.service.GetQuotes(author)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Write(quotes)
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) GetRandomQuote(w http.ResponseWriter, r *http.Request) {
	quote, err := h.service.GetRandomQuote()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Write(quote)
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) DeleteQuoteByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	quoteID, err := strconv.Atoi(idStr)
    if err != nil {
        http.Error(w, "Invalid id parameter", http.StatusBadRequest)
    }
	h.service.DeleteQuoteByID(quoteID)
	w.WriteHeader(http.StatusOK)
}
