package main

import (
	"fmt"
	"net/http"
	"sync"
)

// Custom handler type that includes a counter and a mutex for safe concurrent access
type counterHandler struct {
	count int
	mu    sync.Mutex
}

// ServeHTTP method for counterHandler increments the counter and writes the current count
func (h *counterHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Lock the mutex to ensure safe access to the count variable
	h.mu.Lock()
	h.count++
	count := h.count
	h.mu.Unlock()

	// Respond with the current count
	fmt.Fprintf(w, "This page has been visited %d times", count)
}

func main() {
	// Create an instance of the custom handler
	counter := &counterHandler{}

	counter2 := &counterHandler{}

	// Register the custom handler with a specific route
	http.Handle("/count", counter)

	http.Handle("/counter-2", counter2)

	// Start the server on port 8080
	fmt.Println("Starting server on :8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
