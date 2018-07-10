package main

import (
	"flag"
	"net"
	"fmt"
	"os"
	"context"
)

var host2 string

func main() {
	flag.StringVar(&host, "host", "localhost", "host name to resolve")
	flag.Parse()

	res := net.Resolver{PreferGo: true}
	addrs, err := res.LookupHost(context.Background(), host)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(addrs)
}
