package main

import (
	"flag"
	"net"
	"fmt"
	"os"
)

var host string

func main() {
	flag.StringVar(&host, "host", "localhost", "host name to resolve")
	flag.Parse()

	addrs, err := net.LookupHost(host)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(addrs)
}
