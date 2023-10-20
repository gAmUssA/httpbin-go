package main

import "github.com/gin-gonic/gin"

// DefaultHeaders method returns some default headers like server version
func DefaultHeaders() gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO add condition to display this header
		//like if c.DefaultQuery("MyDebug", "false") == "true"
		c.Header("Server", "Gin "+gin.Version)
		c.Next()
	}
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Next()
	}
}
