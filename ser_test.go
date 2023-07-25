package main

import (
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
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
	SetupJsonRoute(router)

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

func TestXMLRoute(t *testing.T) {
	// Set Gin to Test Mode
	gin.SetMode(gin.TestMode)

	// Setup Routes
	router := gin.Default()
	SetupXmlRoute(router)

	// Create httptest server
	server := httptest.NewServer(router)
	defer server.Close()

	// Make a request to the /xml endpoint
	res, err := http.Get(server.URL + "/xml")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Read the response body
	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Parse the XML response
	var s Slideshow
	err = xml.Unmarshal(body, &s)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Check the XML attributes
	if s.Title != "Sample Slide Show" {
		t.Errorf("Expected title to be 'Sample Slide Show', got %v", s.Title)
	}
	if s.Author != "Yours Truly" {
		t.Errorf("Expected author to be 'Yours Truly', got %v", s.Author)
	}
	if s.Date != "Date of publication" {
		t.Errorf("Expected date to be 'Date of publication', got %v", s.Date)
	}
}
