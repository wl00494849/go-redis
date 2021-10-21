package middleware

import (
	"go-redis/controller"

	"github.com/gin-gonic/gin"
	c "github.com/rs/cors/wrapper/gin"
)

func Setup(app *gin.Engine) {
	app.Use(cors())
	app.Use(choseRedis())
}

func cors() gin.HandlerFunc {
	return c.New(c.Options{
		AllowedOrigins: []string{
			"http://localhost:3000",
			"http://localhost:4200",
			"http://localhost:58505",
			"http://localhost:5000",
		},
	})
}

func choseRedis() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		pattern := ctx.GetHeader("Pattern")
		go controller.TurnPostCache(pattern)
		ctx.Next()
	}
}
