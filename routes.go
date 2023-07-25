package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
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
