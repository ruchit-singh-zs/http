package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"
)

func TestCounterHandler(t *testing.T) {
	tests := []struct {
		name           string
		expectedCounts []int
	}{
		{
			name:           "Single request",
			expectedCounts: []int{1},
		},
		{
			name:           "Multiple sequential requests",
			expectedCounts: []int{1, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			counter := &counterHandler{}
			for i, expectedCount := range tt.expectedCounts {
				req := httptest.NewRequest(http.MethodGet, "http://example.com/count", nil)
				w := httptest.NewRecorder()

				counter.ServeHTTP(w, req)

				res := w.Result()
				body := w.Body.String()

				expectedBody := strings.TrimSpace(fmt.Sprintf("This page has been visited %d times", expectedCount))
				if body != expectedBody {
					t.Errorf("Test %d: Expected body %q, but got %q", i+1, expectedBody, body)
				}

				if res.StatusCode != http.StatusOK {
					t.Errorf("Test %d: Expected status code %d, but got %d", i+1, http.StatusOK, res.StatusCode)
				}
			}
		})
	}
}

func TestCounterHandler_ConcurrentRequests(t *testing.T) {
	const concurrentRequests = 100
	counter := &counterHandler{}
	wg := sync.WaitGroup{}
	wg.Add(concurrentRequests)

	for i := 0; i < concurrentRequests; i++ {
		go func() {
			defer wg.Done()
			req := httptest.NewRequest(http.MethodGet, "http://example.com/count", nil)
			w := httptest.NewRecorder()
			counter.ServeHTTP(w, req)
		}()
	}

	wg.Wait()

	// Verify that the count is as expected after all requests
	expectedCount := concurrentRequests
	if counter.count != expectedCount {
		t.Errorf("Expected count to be %d, but got %d", expectedCount, counter.count)
	}
}
