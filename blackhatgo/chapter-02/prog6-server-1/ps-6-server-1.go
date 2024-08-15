package main

import (
	"io"
	"log"
	"net"
)

func echo(conn net.Conn) {
	defer conn.Close()

	if _, err := io.Copy(conn, conn); err != nil {
		log.Fatalln("Unable to read/write data")
	}

	// reader := bufio.NewReader(conn)

	// s, err := reader.ReadString('\n')
	// if err != nil {
	// 	log.Fatalln("Unable to read data")
	// }
	// log.Printf("Read %d bytes: %s\n", len(s), s)

	// log.Println("Writing data")
	// writer := bufio.NewWriter(conn)
	// if _, err := writer.WriteString(s); err != nil {
	// 	log.Fatalln("Unable to wriye data")
	// }
	// writer.Flush()
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
