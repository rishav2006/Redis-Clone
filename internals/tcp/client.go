package tcp

import (
	"bufio"
	"fmt"
	"net"
	"os"

	"github.com/rishav2006/redis-clone/internals/controllers"
)

func TakeInput() string{
	fmt.Println("Enter the command")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	if input == ""{
		fmt.Println("Error : Please provide some input")
		return ""
	}
	return input
}

func Tester() {
	conn, err := net.Dial("tcp", "localhost:6379")
	if err != nil{
		fmt.Println(err);
		return
	}
	msg := TakeInput()
	conn.Write([]byte(msg))
	controllers.Organizer(msg)
}
