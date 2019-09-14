package main

import (
	"fmt"

	"github.com/cakturk/go-netstat/netstat"
)

var tabs []netstat.SockTabEntry

func displaySocks() error {
	// UDP sockets
	socks, err := netstat.UDPSocks(netstat.NoopFilter)
	if err != nil {
		return err
	}
	for _, e := range socks {
		fmt.Printf("udp: %v\n", e)
	}

	// TCP sockets
	socks, err = netstat.TCPSocks(netstat.NoopFilter)
	if err != nil {
		return err
	}
	for _, e := range socks {
		fmt.Printf("tcp: %v\n", e)
	}

	// get only listening TCP sockets
	tabs, err = netstat.TCPSocks(func(s *netstat.SockTabEntry) bool {
		return s.State == netstat.Listen
	})
	if err != nil {
		return err
	}
	for _, e := range socks {
		fmt.Printf("lst: %v\n", e)
	}

	// list all the TCP sockets in state FIN_WAIT_1 for your HTTP server
	tabs, err = netstat.TCPSocks(func(s *netstat.SockTabEntry) bool {
		return s.State == netstat.FinWait1 && s.LocalAddr.Port == 80
	})
	// error handling, etc.

	return nil
}

func main() {
	displaySocks()
}
