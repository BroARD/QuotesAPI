package service

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"time"
)

//Цитаты
var quotes []*Quote
//Костыль для индексов
var defQuoteID int = 1

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type Service interface {
	CreateQuote(quote Quote) 
	GetQuotes(author string) ([]byte, error)
	GetRandomQuote() ([]byte, error)
	DeleteQuoteByID(quoteID int)
}

type service struct{}

func NewService() Service {
	return &service{}
}

func (s *service) CreateQuote(quote Quote) {
	log.Printf("Add new quote")

	quote.ID = defQuoteID
	defQuoteID++

	quotes = append(quotes, &quote)
}

func (s *service) GetQuotes(author string) ([]byte, error){
	if author != "" {
		jsonQuote, err := json.Marshal(FilterQuotesByAuthor(author))
		if err != nil {
			return nil, err
		}
		return jsonQuote, nil
	} else {
		jsonQuote, err := json.Marshal(quotes)
		if err != nil {
			return nil, err
		}
		return jsonQuote, nil
	}
}

func (s *service) GetRandomQuote() ([]byte, error){
	if len(quotes) == 0 {
		return nil, http.ErrBodyNotAllowed
	}

	rand.Seed(time.Now().Unix())
	jsonData, err := json.Marshal(quotes[rand.Intn(len(quotes))])
	if err != nil {
		return nil, err
	}
	return jsonData, nil
}

func (s *service) DeleteQuoteByID(quoteID int){
	log.Printf("Попытка удалить")
	for index, quote := range quotes {
        if quote.ID == quoteID {
            quotes = append(quotes[:index], quotes[index+1:]...)
        	return
        }
    }
}


func FilterQuotesByAuthor(author string) []*Quote{
	var filterQuotes []*Quote
	for _ , quote := range quotes {
		if quote.Author == author {
			filterQuotes = append(filterQuotes, quote)
		}
	}

	return filterQuotes
}
