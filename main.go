package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type link struct {
	ID       string `json:"id"`
	UrlOrig  string `json:"url"`
	ShortUrl string `json:"ShortUrl"`
	Time     string `json:"Time"`
}

var links = []link{
	{ID: "0",
		UrlOrig:  "https://testlink.com/database/index012345",
		ShortUrl: "https://shortlink.com/0"},
}

func main() {
	router := gin.Default()
	router.GET("/links", getUrlData)
	router.GET("/links/:id", getUrlByID)
	router.POST("/links", postLink)
	router.Run("localhost:8080")
}

func getUrlData(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, links)
}
func postLink(c *gin.Context) {
	var newLink link
	if err := c.BindJSON(&newLink); err != nil {
		return
	}
	links = append(links, newLink)
	c.IndentedJSON(http.StatusCreated, newLink)
}
func getUrlByID(c *gin.Context) {
	id := c.Param("id")
	for _, l := range links {
		if l.ID == id {
			c.IndentedJSON(http.StatusOK, l)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "ID not found"})
}
