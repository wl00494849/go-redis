package main

import (
	"flag"
	"go-redis/controller"
	"go-redis/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	var port string
	flag.StringVar(&port, "Port", ":7788", "server port")

	app := gin.Default()
	//middleware
	middleware.Setup(app)
	//route
	app.POST("/redisPost", controller.SetRedis)

	app.Run(port)
}
