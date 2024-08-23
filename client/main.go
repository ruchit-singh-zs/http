package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	// Create an HTTP client
	client := &http.Client{}

	// Create a GET request
	req, err := http.NewRequest(http.MethodGet, "http://localhost:8080/hello", nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	// Print the response status and body
	fmt.Println("Response status:", resp.Status)
	fmt.Println("Response body:", string(body))
}
