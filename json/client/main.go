package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type RequestData struct {
	Name string `json:"name"`
}

func main() {
	// Create a struct instance with the data to send
	reqData := RequestData{
		Name: "John",
	}

	// Marshal the struct to JSON
	jsonData, err := json.Marshal(reqData)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	// Create an HTTP client and request
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPost, "http://localhost:8080/greet", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Set the content type to application/json
	req.Header.Set("Content-Type", "application/json")

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// Print the response status
	fmt.Println("Response status:", resp.Status)
}
