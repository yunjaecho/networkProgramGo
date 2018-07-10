package main

import (
	"flag"
	"net"
	"fmt"
	"os"
)

var addr string
/**
   ==========================
   The net.Resolver Type
   ==========================
   Type Resolver exposes several methods to help resolve names or query other
   DSN information of attached hosts to the network.
   The Resolver can use either a pure GO implementation, which make direct DNS Calls from GO
   Or, it can make C library system calls using Cgo
 */
func main() {
	flag.StringVar(&addr, "addr", "127.0.0.1", "host address to lookup")
	flag.Parse()

	name, err := net.LookupAddr(addr)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(name)
}
