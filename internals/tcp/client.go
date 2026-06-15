package tcp

import (
	"fmt"
	"net"
)

func Tester() {
	conn, err := net.Dial("tcp", "localhost:6379")
	if err != nil{
		fmt.Println(err);
		return
	}
	conn.Write([]byte("hello"))
}
