package tcp

import (
	"fmt"
	"net"
)

func TcpConnect() {
	listener, err := net.Listen("tcp", ":6379")
	if err != nil {
		fmt.Println("An error occured during listener process : ", err)
		return
	}
	fmt.Println("Server started")
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("An error occured while accepting new connections : ", err)
			continue
		}

		buffer := make([]byte, 1024)

		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Println("Received:", string(buffer[:n]))
	}
}
