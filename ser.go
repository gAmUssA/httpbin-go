package main

import (
	"fmt"
	"math/rand"
	"net"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type IPResponse struct {
	Origin string `json:"origin"`
}

type UUIDResponse struct {
	UUID string `json:"uuid"`
}

type UserAgentResponse struct {
	UserAgent string `json:"user-agent"`
}

type GetResponse struct {
	Args    map[string]string `json:"args"`
	Headers map[string]string `json:"headers"`
	Url     string            `json:"url"`
}

type PostResponse struct {
	Args    map[string]string `json:"args"`
	Data    string            `json:"data"`
	Headers map[string]string `json:"headers"`
	Json    gin.H             `json:"json"`
	Url     string            `json:"url"`
}
type AnythingResponse struct {
	Method  string            `json:"method"`
	Headers map[string]string `json:"headers"`
	Url     string            `json:"url"`
}

func main() {
	router := gin.Default()

	router.GET("/ip", func(c *gin.Context) {
		ip, _, _ := net.SplitHostPort(c.Request.RemoteAddr)
		c.JSON(http.StatusOK, IPResponse{Origin: ip})
	})

	router.GET("/uuid", func(c *gin.Context) {
		uuid := generateUUID()
		c.JSON(http.StatusOK, UUIDResponse{UUID: uuid})
	})

	router.GET("/user-agent", func(c *gin.Context) {
		c.JSON(http.StatusOK, UserAgentResponse{UserAgent: c.Request.UserAgent()})
	})

	router.GET("/get", func(c *gin.Context) {
		headers, args := make(map[string]string), make(map[string]string)
		for name, values := range c.Request.Header {
			headers[name] = values[0]
		}
		for name, values := range c.Request.URL.Query() {
			args[name] = values[0]
		}
		c.JSON(http.StatusOK, GetResponse{
			Args:    args,
			Headers: headers,
			Url:     c.Request.RequestURI,
		})
	})

	router.POST("/post", func(c *gin.Context) {
		var jsonBody gin.H
		c.BindJSON(&jsonBody)
		headers, args := make(map[string]string), make(map[string]string)
		for name, values := range c.Request.Header {
			headers[name] = values[0]
		}
		for name, values := range c.Request.URL.Query() {
			args[name] = values[0]
		}
		c.JSON(http.StatusOK, PostResponse{
			Args:    args,
			Data:    c.Request.Form.Encode(),
			Headers: headers,
			Json:    jsonBody,
			Url:     c.Request.RequestURI,
		})
	})

	router.GET("/status/:status_code", func(c *gin.Context) {
		statusCode, _ := strconv.Atoi(c.Param("status_code"))
		c.JSON(statusCode, gin.H{"message": http.StatusText(statusCode)})
	})

	router.GET("/anything", func(c *gin.Context) {
		headers := make(map[string]string)
		for name, values := range c.Request.Header {
			headers[name] = values[0] // Only get the first value of each header
		}
		c.JSON(http.StatusOK, AnythingResponse{
			Method:  c.Request.Method,
			Headers: headers,
			Url:     c.Request.RequestURI,
		})
	})

	router.Run(":8080")
}

func generateUUID() string {
	b := make([]byte, 16)
	rand.Read(b)
	return fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
}
