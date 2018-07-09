package main

import (
	"os"
	"fmt"
	"net"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Missing ip address")
		os.Exit(1)
	}

	ip := net.ParseIP(os.Args[1])
	if ip == nil {
		fmt.Println("Unable to parse IP Address")
		fmt.Println("Address should use IPv4 dot-notation or IPv6 colon-notation")
		os.Exit(1)
	}

	fmt.Println()
	fmt.Printf("IP:					%s\n", ip)
	fmt.Printf("Default Mask: 		%s\n", net.IP(ip.DefaultMask()))
	fmt.Printf("Loopback:			%t\n", ip.IsLoopback())
	fmt.Println("Unicast:")
	fmt.Printf("  Global:			%t\n", ip.IsGlobalUnicast())
	fmt.Printf("  Link:				%t\n", ip.IsLinkLocalUnicast())
	fmt.Println("Multicast:")
	fmt.Printf("  Global: 			%t\n", ip.IsMulticast())
	fmt.Printf("  Interface			%t\n", ip.IsInterfaceLocalMulticast())
	fmt.Printf("  Link				%t\n", ip.IsLinkLocalMulticast())
	fmt.Println()
}
