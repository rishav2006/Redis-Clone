package tcp

import (
	"fmt"
	"net"

	"github.com/rishav2006/redis-clone/internals/controllers"
)

func HandleConnection(conn net.Conn) {
	defer conn.Close()

	for {
		buffer := make([]byte, 1024)

		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Client disconnected")
			return
		}

		fmt.Println("Received:", string(buffer[:n]))
		msg := string(buffer[:n])
		answer := controllers.Organizer(msg, conn)
		conn.Write([]byte(answer))
	}
}

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
		fmt.Println("New Client Connected")
		go HandleConnection(conn)
	}
}
