package main

import (
	// "github.com/rishav2006/redis-clone/internals/controllers"
	"github.com/rishav2006/redis-clone/internals/persistance"
	"github.com/rishav2006/redis-clone/internals/tcp"
)

// "fmt"
// "github.com/rishav2006/redis-clone/internals/routes"

func main() {
	
	// r := routes.NewRouter()
	// r.Run(":6379")
	// fmt.Println("Server Running on PORT:6379")
	// for 1 < 2{
	// 	controllers.Organizer();
	// }
	
	persistance.LoadSnapshot()
	tcp.TcpConnect();
}
