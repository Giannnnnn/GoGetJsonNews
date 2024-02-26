package main

import (
	"encoding/xml"
	"net/http"
	"github.com/gin-gonic/gin"
)

type Entry struct {
	ID         string     `xml:"id"`
	Title      string     `xml:"title"`
	Summary    string     `xml:"summary"`
	Published  string     `xml:"published"`
	Link       Link       `xml:"link"`
	Categories []Category `xml:"category"`
}

type Link struct {
	Rel  string `xml:"rel,attr"`
	Href string `xml:"href,attr"`
}

type Category struct {
	Term string `xml:"term,attr"`
}

type Feed struct {
	Entries []Entry `xml:"entry"`
}

func fetchRSSFeed(url string) (Feed, error) {
	resp, err := http.Get(url)
	if err != nil {
		return Feed{}, err
	}
	defer resp.Body.Close()

	var feed Feed
	err = xml.NewDecoder(resp.Body).Decode(&feed)
	if err != nil {
		return Feed{}, err
	}

	return feed, nil
}

func topicHandler(c *gin.Context) {
	platform := c.Param("platform")
	topic := c.Param("topic")

	var url string
	switch platform {
	case "g1":
		url = g1Topics[topic]
	case "bbc":
		url = bbcTopics[topic]
	default:
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Platform not found",
		})
		return
	}

	if url == "" {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Topic not found",
		})
		return
	}

	feed, err := fetchRSSFeed(url)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch RSS feed",
		})
		return
	}

	c.JSON(http.StatusOK, feed)
}

var g1Topics = map[string]string{
	"brasil":              "https://g1.globo.com/dynamo/brasil/rss2.xml",
	"carros":              "https://g1.globo.com/dynamo/carros/rss2.xml",
	"ciencia-e-saude":     "https://g1.globo.com/dynamo/ciencia-e-saude/rss2.xml",
	"concursos-e-emprego": "https://g1.globo.com/dynamo/concursos-e-emprego/rss2.xml",
}

var bbcTopics = map[string]string{
	"brasil":         "http://www.bbc.co.uk/portuguese/topicos/brasil/index.xml",
	"america_latina": "http://www.bbc.co.uk/portuguese/topicos/america_latina/index.xml",
	"internacional":  "http://www.bbc.co.uk/portuguese/topicos/internacional/index.xml",
}

func main() {
	r := gin.Default()

	r.GET("/topics/:platform/:topic", topicHandler)

	r.Run(":8080")
}
