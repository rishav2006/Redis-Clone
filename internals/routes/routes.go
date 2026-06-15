package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rishav2006/redis-clone/internals/controllers"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	r.POST("/post", controllers.Organizer())
	return r
}