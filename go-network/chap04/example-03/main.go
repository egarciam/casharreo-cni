package main

//TCP SERVER ECHO

import (
	"fmt"
	"net"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()
	buffer := make([]byte, 1024)
	for {
		n, err := conn.Read(buffer)
		fmt.Println("Leido: ", buffer[:n])
		if err != nil {
			return
		}
		_, err = conn.Write(buffer[:n])
		if err != nil {
			return
		}
	}
}

func main() {
	listener, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Println("error al crear listener")
		return
	}

	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("error en Accept")
		}

		go handleConnection(conn)
	}
}
