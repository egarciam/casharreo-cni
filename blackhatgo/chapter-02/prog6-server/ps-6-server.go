package main

import (
	"io"
	"log"
	"net"
)

func echo(conn net.Conn) {
	defer conn.Close()

	b := make([]byte, 512)

	for {
		size, err := conn.Read(b[0:])
		if err == io.EOF {
			log.Println("Client disconnected")
			break
		}
		if err != nil {
			log.Println("Unexpected error")
			break
		}
		log.Printf("Received %d bytes: %s\n", size, string(b))

		//Send data via conn.Write

		log.Println("Writing data")

		if _, err := conn.Write(b[0:size]); err != nil {
			log.Fatalln("Unable to write data")
		}
	}
}

func main() {
	//Bind to port 20080 on all ifaces
	listener, err := net.Listen("tcp", ":20080")

	if err != nil {
		log.Fatalln("Unable to bind port")
	}

	log.Println("Listening on 0.0.0.0:20080")

	for {
		conn, err := listener.Accept()
		log.Println("Received connection")

		if err != nil {
			log.Fatalln("unable to accept connection")
		}

		//handle the connection
		go echo(conn)
	}
}
