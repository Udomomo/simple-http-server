package main

import (
	"fmt"
	"net"
	"os"
)

const PORT = 8000

func main() {
	// Socketを作成し、IPアドレス・ポートにbindする
	address := fmt.Sprintf("localhost:%d", PORT)

	tcpAdrr, err := net.ResolveTCPAddr("tcp4", address)
	if err != nil {
		fmt.Printf("Error: %s", err)
		os.Exit(1)
	}
	listener, err := net.ListenTCP("tcp", tcpAdrr)
	if err != nil {
		fmt.Printf("Error: %s", err)
		os.Exit(1)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("Error: %s", err)
			continue
		}

		request, err := Parse(conn)
		if err == nil {
			fmt.Printf("Request is parsed. %#v\n", request)
		}
	}
}
