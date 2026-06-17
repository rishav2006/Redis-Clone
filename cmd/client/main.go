package main

import (
	"fmt"
	"net"

	"github.com/rishav2006/redis-clone/internals/tcp"
)

func ListenForMessages(conn net.Conn) {
	for {
		buffer := make([]byte, 1024)

		n, err := conn.Read(buffer)

		if err != nil {
			fmt.Println("\nConnection has been closed")
			return
		}

		fmt.Printf("\nMESSAGE: %s\n", string(buffer[:n]))
	}
}

func main() {
	conn, err := net.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println(err)
		return
	}

	go ListenForMessages(conn)

	for {
		msg := tcp.TakeInput()

		_, err := conn.Write([]byte(msg))
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
