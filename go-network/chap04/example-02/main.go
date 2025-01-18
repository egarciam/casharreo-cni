package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	// Connect to a remote server (example.com) on port 80 (HTTP)
	addrs, err := net.LookupHost("google.com")

	if err != nil {
		fmt.Println("Error resolving hostname:", err)
		return
	}

	for _, addr := range addrs {
		fmt.Println(addr)
	}

	addr, err := net.ResolveUDPAddr("udp4", ":8991")
	if err != nil {
		fmt.Println("Error resolving UDP ADDR:", err)
		return
	}
	fmt.Println(addr)
	conn, err := net.ListenUDP("udp4", addr)
	if err != nil {
		panic(1)
	} else {
		fmt.Println("udp listening")
	}

	go func() {
		_, err := net.Dial("tcp4", "8.8.8.8:53")
		if err != nil {
			fmt.Println("error in dial")
			return
		}
		fmt.Println("internet connection active")
	}()

	time.Sleep(300 * time.Second)
	defer conn.Close()

}
