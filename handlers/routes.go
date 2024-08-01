package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(CORSMiddleware())
	r.Use(LoggingMiddleware())

	r.LoadHTMLGlob("handlers/views/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home.html", nil)
	})

	r.GET("/customers", GetCustomers)
	r.GET("/customers/:id", GetCustomer)
	r.POST("/customers", CreateCustomer)
	r.PUT("/customers/:id", UpdateCustomer)
	r.DELETE("/customers/:id", DeleteCustomer)

	return r
}
