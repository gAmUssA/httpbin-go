package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestJsonEndpoint(t *testing.T) {
	// Set Gin to Test Mode
	gin.SetMode(gin.TestMode)

	// Setup Routes
	router := gin.Default()
	setupJsonRoute(router)

	// Create httptest server
	server := httptest.NewServer(router)
	defer server.Close()

	// Create a request to our server with the {base url}/json
	request, err := http.NewRequest(http.MethodGet, server.URL+"/json", nil)

	if err != nil {
		t.Fatalf("Couldn't create request: %s\n", err)
	}

	response, err := http.DefaultClient.Do(request)

	if err != nil {
		t.Fatalf("Couldn't send request: %s\n", err)
	}

	// We should have a 200 status
	assert.Equal(t, http.StatusOK, response.StatusCode)

	// Decode the JSON response
	var jsonData JSONSlideshow
	err = json.NewDecoder(response.Body).Decode(&jsonData)

	if err != nil {
		t.Fatalf("Couldn't decode json response: %s\n", err)
	}

	// Assert the jsonData is as expected
	expectedData := JSONSlideshow{
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
	assert.Equal(t, expectedData, jsonData)
}
