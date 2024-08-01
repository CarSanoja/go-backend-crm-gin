package handlers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	return cors.New(config)
}

func LoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}
