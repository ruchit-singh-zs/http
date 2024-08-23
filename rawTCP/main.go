package main

import (
	"bufio"
	"fmt"
	"net"
	"net/http"
	"strings"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Server is listening on :8080...")

	for {
		// A Listener is a generic network listener for stream-oriented protocols.
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)
	requestLine, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading request:", err)
		return
	}

	// Very basic request parsing
	method := strings.Fields(requestLine)[0]
	path := strings.Fields(requestLine)[1]

	if method == http.MethodGet && path == "/hello" {
		response := "HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\n\r\nHello, World!"
		conn.Write([]byte(response))
	} else {
		response := "HTTP/1.1 404 Not Found\r\nContent-Type: text/plain\r\n\r\nNot Found"
		conn.Write([]byte(response))
	}
}
