package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func setupJsonRoute(router *gin.Engine) {
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
