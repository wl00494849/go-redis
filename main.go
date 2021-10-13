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
	app.GET("/redisGet", controller.GetRedis)
	app.POST("/redisPush", controller.PushRedis)
	app.GET("/redisLrange", controller.LrangeRedis)

	app.Run(port)
}
