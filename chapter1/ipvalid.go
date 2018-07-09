package main

import (
	"os"
	"net"
	"fmt"
)

/**
	IP Address Parse
 */
func main() {
	if len(os.Args) != 2 {
		return
	}

	// parse IP address
	// The string s can be in dotted decimal("192.0.2.1") or
	// IPv6 ("2001"db8::68") from
	ip := net.ParseIP(os.Args[1])

	if ip != nil {
		fmt.Printf("%v OK\n", ip)
	} else {
		fmt.Println("Bad address")
	}
}
