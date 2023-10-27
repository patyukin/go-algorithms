package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("Starting the server...")
	listener, err := net.Listen("tcp", "0.0.0.0:3001")
	if err != nil {
		fmt.Println("Error listening", err.Error())
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			return
		}

		go doServerStuff(conn)
	}
}

func doServerStuff(conn net.Conn) {
	for {
		buf := make([]byte, 512)
		_, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error reading", err.Error())
			return
		}

		fmt.Printf("Received data: %v\n", string(buf))
	}
}
