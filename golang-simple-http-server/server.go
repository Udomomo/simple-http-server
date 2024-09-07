package main

import (
	"fmt"
	"net"
	"os"
)

const PORT = 8080

func main() {
	// Socketを作成し、IPアドレス・ポートにbindする
	tcpAdrr, err := net.ResolveTCPAddr("tcp4", "localhost:8080")
	if (err != nil) {
		fmt.Printf("Error: %s", err)
		os.Exit(1)
	}
	listener, err := net.ListenTCP("tcp", tcpAdrr)
	if (err != nil) {
		fmt.Printf("Error: %s", err)
		os.Exit(1)
	}

	for {
		conn, err := listener.Accept()
		if (err != nil) {
			fmt.Printf("Error: %s", err)
			continue
		}
	}
}
