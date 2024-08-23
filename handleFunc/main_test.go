package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHomeHandler(t *testing.T) {
	tests := []struct {
		name           string
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "Home route",
			expectedStatus: http.StatusOK,
			expectedBody:   "Welcome to the Go HTTP Server!",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "http://example.com/", nil)
			w := httptest.NewRecorder()

			homeHandler(w, req)

			res := w.Result()
			if res.StatusCode != tt.expectedStatus {
				t.Errorf("Expected status code %d, but got %d", tt.expectedStatus, res.StatusCode)
			}

			if body := w.Body.String(); body != tt.expectedBody {
				t.Errorf("Expected body %q, but got %q", tt.expectedBody, body)
			}
		})
	}
}

func TestHelloHandler(t *testing.T) {
	tests := []struct {
		name           string
		query          string
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "Hello with no name parameter",
			query:          "",
			expectedStatus: http.StatusOK,
			expectedBody:   "Hello, World!",
		},
		{
			name:           "Hello with name parameter",
			query:          "name=Alice",
			expectedStatus: http.StatusOK,
			expectedBody:   "Hello, Alice!",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "http://example.com/hello?"+tt.query, nil)
			w := httptest.NewRecorder()

			helloHandler(w, req)

			res := w.Result()
			if res.StatusCode != tt.expectedStatus {
				t.Errorf("Expected status code %d, but got %d", tt.expectedStatus, res.StatusCode)
			}

			if body := w.Body.String(); body != tt.expectedBody {
				t.Errorf("Expected body %q, but got %q", tt.expectedBody, body)
			}
		})
	}
}
