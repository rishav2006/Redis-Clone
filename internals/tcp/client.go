package tcp

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func TakeInput() string {
	fmt.Println("Enter the command")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	if input == "" {
		fmt.Println("Error : Please provide some input")
		return ""
	}
	return input
}

func Tester() {
	conn, err := net.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println(err)
		return
	}
	msg := TakeInput()
	conn.Write([]byte(msg))
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	fmt.Println(string(buffer[:n]))
}
