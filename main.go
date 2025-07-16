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

func getUrlData(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, links)
}

func main() {
	router := gin.Default()
	router.GET("/links", getUrlData)
	router.Run("localhost:8080")
}
