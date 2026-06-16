package main

import (
	"fmt"
	"net"

	"github.com/rishav2006/redis-clone/internals/tcp"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		msg := tcp.TakeInput()
		conn.Write([]byte(msg))
		buffer := make([]byte, 1024)
		n, _ := conn.Read(buffer)
		fmt.Println(string(buffer[:n]))
	}
}
