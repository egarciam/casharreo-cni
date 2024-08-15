package main

import (
	"fmt"
	"net"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	for i := 1; i <= 65535; i++ {
		wg.Add(1)
		go func(j int) {
			defer wg.Done()
			//address := fmt.Sprintf("scanme.nmap.org:%d", i)
			address := fmt.Sprintf("127.0.0.1:%d", i)
			conn, err := net.Dial("tcp", address)
			if err != nil {
				//sfmt.Println("Err: ", err)
				return
			}
			conn.Close()
			fmt.Printf("%d open\n", i)
		}(i)

	}
	wg.Wait()
}
