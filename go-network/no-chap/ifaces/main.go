package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

// This program lists network interface information.
func main() {
	var ifname string
	if len(os.Args) == 2 {
		ifname = os.Args[1]
	}
	ifaces, err := net.Interfaces()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, iface := range ifaces {
		if ifname == "" || ifname == iface.Name {
			addrs, err := iface.Addrs()
			if err != nil {
				continue
			}
			fmt.Printf("%s: <%s> mtu=%d\n", iface.Name, strings.ToUpper(iface.Flags.String()), iface.MTU)
			if len(iface.HardwareAddr.String()) > 0 {
				fmt.Printf("\tether %s\n", iface.HardwareAddr.String())
			}
			if len(addrs) > 0 {
				for _, addr := range addrs {
					//fmt.Printf("\t%s\n", addrInfo(addr))
					fmt.Printf("\t%v\n", addr)
				}
			}
		}
	}
}
