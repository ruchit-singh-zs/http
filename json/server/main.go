package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Struct for the JSON request
type RequestData struct {
	Name string `json:"name"`
}

// Struct for the JSON response
type ResponseData struct {
	Greeting string `json:"greeting"`
}

func main() {
	// Set up a route that handles the /greet endpoint
	http.HandleFunc("/greet", greetHandler)

	// Start the server on port 8080
	fmt.Println("Server is running on http://localhost:8080...")
	http.ListenAndServe(":8080", nil)
}

// The handler function for /greet
func greetHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	// Decode the JSON request into RequestData struct
	var reqData RequestData

	// OPTIONAL :
	//try out with file

	err := json.NewDecoder(r.Body).Decode(&reqData)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Create a response with a greeting message
	respData := ResponseData{
		Greeting: "Hello, " + reqData.Name + "!",
	}

	// Encode the response as JSON and send it back to the client
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(respData)
}
