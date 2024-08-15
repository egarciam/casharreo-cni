package main

import "net"

func main() {
	_, err := net.Dial("tcp", "scanme.nmap.org:80")
	if err == nil {
		println("Connection successful")
	} else {
		println("Failed")
	}

	return
}
