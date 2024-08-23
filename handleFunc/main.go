package main

import (
	"fmt"
	"log"
	"net/http"
)

// Handler function for the root route "/"
func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		fmt.Fprintf(w, "Welcome to the Go HTTP Server!")
	} else {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("route not found"))
		// fmt.Fprintf(w, "route not found")
	}
}

// Handler function for the "/hello" route
func helloHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "World"
	}

	fmt.Fprintf(w, "Hello, %s!", name)
}

func main() {
	// Register handler functions for specific routes
	http.HandleFunc("/", homeHandler)
	// http.HandleFunc("/hello", homeHandler)
	http.HandleFunc("/hello", helloHandler)

	// Start the server on port 8080
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err.Error())
	}
}
