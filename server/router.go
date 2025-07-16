package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/links", getUrlData)
	router.GET("/links/:id", getUrlByID)
	router.POST("/links", postLink)
	return router
}

func getUrlData(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, links)
}
func postLink(c *gin.Context) {
	var newLink Link
	if err := c.BindJSON(&newLink); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, nil)
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
