package handlers

import (
	"APIWithout/internal/service"
	mock_service "APIWithout/internal/service/mocks"
	"bytes"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestHandler_PostQuote(t *testing.T) {
	type mockBehavior func(s *mock_service.MockService, quote service.Quote)

	testTable := []struct {
		name         string
		inputBody    string
		inputQuote   service.Quote
		mockBehavior mockBehavior
		expectedStatusCode int
		expectedRequestBody string
	} {
		{
			name: "OK",
			inputBody: `{"author":"Test", "quote":"qwerty"}`,
			inputQuote: service.Quote{
				Author: "Test",
				Quote: "qwerty",
			},
			mockBehavior: func(s *mock_service.MockService, quote service.Quote) {
				s.EXPECT().CreateQuote(quote).Return(1, nil)
			},
			expectedStatusCode: 200,
			expectedRequestBody: "",
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()
			
			mock := mock_service.NewMockService(c)
			testCase.mockBehavior(mock, testCase.inputQuote)

			service := service.NewService()
			handlers := NewHandler(service)

			r := mux.NewRouter()
			r.HandleFunc("/quotes", handlers.PostQuote).Methods("POST")

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/quotes", bytes.NewBufferString(testCase.inputBody))

			r.ServeHTTP(w, req)

			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedRequestBody, w.Body.String())
		})
	}
}
