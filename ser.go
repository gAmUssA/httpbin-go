package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jaswdr/faker"
	"net"
	"net/http"
	"strconv"
	"strings"
)

func setupRouter(router *gin.Engine) {
	router.Use(Cors())

	router.GET("/ip", func(c *gin.Context) {
		ip, _, _ := net.SplitHostPort(c.Request.RemoteAddr)
		c.JSON(http.StatusOK, IPResponse{Origin: ip})
	})

	router.GET("/uuid", func(c *gin.Context) {
		uuid := generateUUID()
		c.JSON(http.StatusOK, UUIDResponse{UUID: uuid})
	})

	router.GET("/base64/:value", func(c *gin.Context) {
		encodedValue := c.Param("value")
		decodedBytes, err := base64.URLEncoding.DecodeString(encodedValue)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, string(decodedBytes))
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
		err := c.BindJSON(&jsonBody)
		if err != nil {
			return
		}
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

	router.GET("/response-headers", func(c *gin.Context) {
		queryParams := c.Request.URL.Query()
		result := make(map[string]string)
		for key, values := range queryParams {
			// If multiple values for the same key are present, join them with a comma
			//c.Writer.Header().Set(key, strings.Join(values, ","))

			result[key] = strings.Join(values, ",")
		}
		// append to existing response headers
		header := c.Writer.Header()
		for key, values := range header {
			result[key] = strings.Join(values, ",")
		}
		c.JSON(http.StatusOK, result)
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

	router.GET("/html", func(c *gin.Context) {
		fake := faker.New()

		var text = fake.Lorem().Paragraph(1)
		htmlContent := `
<!DOCTYPE html>
<html>
  <head>
    Hello Random
  </head>
  <body>
    <h1>Lorem Ipsum</h1>
    <div>
      <p>` + text +
			`</p>
    </div>
  </body>
</html>
`
		c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(htmlContent))
	})

	router.GET("/xml", func(c *gin.Context) {
		slideshow := Slideshow{
			Title:  "Sample Slide Show",
			Date:   "Date of publication",
			Author: "Yours Truly",
			Slide: []Slide{
				{
					Type:  "all",
					Title: "Wake up to WonderWidgets!",
				},
				{
					Type:  "all",
					Title: "Overview",
					Item: []Item{
						{Content: "Why <em>WonderWidgets</em> are great"},
						{Content: ""},
						{Content: "Who <em>buys</em> WonderWidgets"},
					},
				},
			},
		}
		c.XML(http.StatusOK, slideshow)
	})

	setupJsonRoute(router)
}

func main() {
	router := gin.Default()
	setupRouter(router)
	err := router.Run(":8080")
	if err != nil {
		return
	}
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Next()
	}
}

func generateUUID() string {
	b := make([]byte, 16)
	_, _ = rand.Read(b)
	return fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
}
