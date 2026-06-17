package tcp

import (
	"bufio"
	"fmt"
	"os"
)

func TakeInput() string {
	fmt.Print("> ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	if input == "" {
		fmt.Println("Error : Please provide some input")
		return ""
	}
	return input
}

// func ListenForMessages(conn net.Conn) {
// 	for {
// 		buffer := make([]byte, 1024)

// 		n, err := conn.Read(buffer)

// 		if err != nil {
// 			fmt.Println("Connection closed")
// 			return
// 		}

// 		fmt.Println(string(buffer[:n]))
// 	}
// }

// func Tester() {
// 	conn, err := net.Dial("tcp", "localhost:6379")
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	go ListenForMessages(conn)
// 	for {
// 		msg := TakeInput()
// 		conn.Write([]byte(msg))
// 	}
// 	// buffer := make([]byte, 1024)
// 	// n, err := conn.Read(buffer)
// 	// fmt.Println(string(buffer[:n]))
// }
