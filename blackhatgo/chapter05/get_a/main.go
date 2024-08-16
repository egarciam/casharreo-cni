package main

import (
	"fmt"

	"github.com/miekg/dns"
)

func main() {
	var msg dns.Msg
	fqdn := dns.Fqdn("stacktitan.com")
	msg.SetQuestion(fqdn, dns.TypeA)
	in, err := dns.Exchange(&msg, "172.27.224.1:53")
	if err != nil {
		panic(err)
	}
	if len(in.Answer) < 1 {
		fmt.Println("no records")
		return
	}

	for _, answer := range in.Answer {
		if a, ok := answer.(*dns.A); ok {
			fmt.Print(a.A)
		}
	}
}
