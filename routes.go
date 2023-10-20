package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func SetupXmlRoute(router *gin.Engine) {
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

}

func SetupJsonRoute(router *gin.Engine) {
	router.GET("/json", func(c *gin.Context) {
		slideshow := JSONSlideshow{
			Title:  "Sample Slide Show",
			Date:   "date of publication",
			Author: "Yours Truly",
			Slides: []JSONSlide{
				{
					Type:  "all",
					Title: "Wake up to WonderWidgets!",
				},
				{
					Type:  "all",
					Title: "Overview",
					Items: []string{
						"Why <em>WonderWidgets</em> are great",
						"Who <em>buys</em> WonderWidgets",
					},
				},
			},
		}
		c.JSON(http.StatusOK, slideshow)
	})
}

func SetupRootRoute(router *gin.Engine) {
	router.GET("/", func(c *gin.Context) {
		// Get all registered routes
		routes := router.Routes()

		var builder strings.Builder
		builder.WriteString("<h1>Welcome to My httpcan</h1>")
		builder.WriteString("<h2>List of all endpoints:</h2>")
		builder.WriteString("<ul>")
		for _, route := range routes {
			builder.WriteString(fmt.Sprintf("<li><strong>%s</strong> - <a href=\"%s\">%s</a></li>", route.Method, stripLeadingSlash(route.Path), route.Path))
		}
		builder.WriteString("</ul>")

		c.Header("Content-Type", "text/html")
		c.String(http.StatusOK, builder.String())
	})
}
