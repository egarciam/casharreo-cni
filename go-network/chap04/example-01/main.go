package main

import (
	"fmt"
	"net"
)

func main() {
	// Connect to a remote server (example.com) on port 80 (HTTP)
	conn, err := net.Dial("tcp", "example.com:80")
	if err != nil {
		fmt.Println("Error connecting:", err)
		return
	}
	defer conn.Close()

	// Send data to the server (GET request)

	_, err = conn.Write([]byte("GET / HTTP/1.1\r\nHost: example.com\r\n\r\n"))
	if err != nil {
		fmt.Println("Error sending data:", err)
		return
	}
	// Receive data from the server (response)
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error receiving data:", err)
		return
	}
	// Print the received data (part of the response)
	fmt.Println(string(buffer[:n]))
}
