package middleware

import (
	"github.com/gin-gonic/gin"
	c "github.com/rs/cors/wrapper/gin"
)

func Setup(app *gin.Engine) {
	app.Use(cors())
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
