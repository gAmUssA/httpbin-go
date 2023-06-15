package main

import (
	"encoding/xml"
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

type Slideshow struct {
	XMLName xml.Name `xml:"slideshow"`
	Title   string   `xml:"title,attr"`
	Date    string   `xml:"date,attr"`
	Author  string   `xml:"author,attr"`
	Slide   []Slide  `xml:"slide"`
}

type Slide struct {
	Type  string `xml:"type,attr"`
	Title string `xml:"title"`
	Item  []Item `xml:"item"`
}

type Item struct {
	Content string `xml:",innerxml"`
}

type JSONSlideshow struct {
	Author string      `json:"author"`
	Date   string      `json:"date"`
	Slides []JSONSlide `json:"slides"`
	Title  string      `json:"title"`
}

type JSONSlide struct {
	Title string   `json:"title"`
	Type  string   `json:"type"`
	Items []string `json:"items,omitempty"`
}
