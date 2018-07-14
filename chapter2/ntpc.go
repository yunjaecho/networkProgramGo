package main

import (
	"flag"
	"net"
	"fmt"
	"os"
	"encoding/binary"
	"time"
)

/**
	Datagram Socket Programming with UDP
    UDP Client
    Usage :
       ntpc -e <host endpoint>
 */
func main() {
	var host string
	flag.StringVar(&host, "e", "us.pool.ntp.org:123", "NTP host")
	flag.Parse()
	// different server
	// go run ntpc.go -e time.nist.gov:123

	// req datagram is a 49-byte long slice
	// the is used for sending time request to the server
	req := make([]byte, 48)

	// req is initialized with 0x1B or 0001 1011 which is
	// a request setting for time server
	// See spec at ntp.org
	req[0] = 0x1B

	// response 48-byte long slice incoming datagram
	// with time values from the server
	rsp := make([]byte, 48)

	// create an address of type UDPAddr that represents
	// the remote host endpoint
	raddr, err := net.ResolveUDPAddr("udp", host)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// setup connection (net.UDPConn) with net.DialUDP()
	conn, err := net.DialUDP("udp", nil, raddr)
	if err != nil {
		fmt.Printf("failed to connect: %v\n", err)
		os.Exit(1)
	}
	defer func() {
		if err:= conn.Close(); err != nil {
			fmt.Println("failed while closing connection: ", err)
		}
	}()

	// send time request
	if _, err := conn.Write(req); err != nil {
		fmt.Printf("failed to send request: %v\n", err)
		os.Exit(1)
	}

	// block to receive server response
	read, err := conn.Read(rsp)
	if err != nil {
		fmt.Printf("failed to receive response: %v\n", err)
		os.Exit(1)
	}

	// ensure we read 48 bytes back (NTP protocol spec)

	if read != 48 {
		fmt.Println("did not get all expected bytes from server")
		os.Exit(1)
	}

	// NTP data comes in as big-endian(LSB [0...47] MSB)
	// with a 64-bit value containing the server time in seconds
	// where the first 32-bits are seconds and last 32-bit are fractional
	// The following extracts the seconds from [0...[40:43]...47]
	// it is the number of secs since 1900 (NTP epoch)
	secs := binary.BigEndian.Uint32(rsp[40:])
	frac := binary.BigEndian.Uint32(rsp[44:])

	ntpEpoch := time.Date(1900, 1, 1, 0, 0, 0, 0, time.UTC)
	unixEpoch := time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC)
	offset := unixEpoch.Sub(ntpEpoch).Seconds()
	now := float64(secs) - offset
	fmt.Printf("%v\n", time.Unix(int64(now), int64(frac)))
}

/*func getNTPSeconds(t time.Time) (int64, int64) {
	secs := t.Unix() + int64(getNTPOffset())
	fracs := t.Nanosecond()
	return secs, int64(fracs)
}
func getNTPOffset() float64 {
	ntpEpoch := time.Date(1900, 1, 1, 0, 0, 0, 0, time.UTC)
	unixEpoch := time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC)
	offset := unixEpoch.Sub(ntpEpoch).Seconds()
	return offset
}


*/
