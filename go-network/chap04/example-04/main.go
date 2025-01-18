package main

import (
	"fmt"
	"net"
)

func handleConnection(conn *net.UDPConn, addr *net.UDPAddr, buffer []byte, n int) {
	// Echo the received data back to the client
	fmt.Printf("Received: %s from %s\n", string(buffer[:n]), addr)
	_, err := conn.WriteToUDP(buffer[:n], addr)
	if err != nil {
		fmt.Println("Error sending response:", err)
	}
}

func main() {
	addr, err := net.ResolveUDPAddr("udp", ":9999")
	if err != nil {
		fmt.Println("Error resolving address:", err)
		return
	}

	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Println("Error creating listener:", err)
		return
	}
	defer conn.Close()

	buffer := make([]byte, 1024)
	for {
		n, addr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println("Error reading from connection:", err)
			continue
		}
		go handleConnection(conn, addr, buffer, n)
	}
}
