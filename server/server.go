package server

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/deatil/go-encoding/base62"
	"github.com/gin-gonic/gin"
)

var lastID int = len(links) + 1

func SetupRouter() *gin.Engine {
	// gin.SetMode(gin.ReleaseMode)
	//TODO: **add reading last freeID on init
	router := gin.Default()
	router.GET("/links", getUrlData)
	router.GET("/links/:id", getUrlByID)
	router.GET("/links/total", getTotalID)
	// router.GET("/links/latest", getLatestID)
	router.POST("/send", postLink)
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
	newLink.ID = int(lastID)
	newLink.ShortID = convertIDBase62(newLink.ID)
	links = append(links, newLink)
	c.IndentedJSON(http.StatusCreated, newLink)
}
func getUrlByID(c *gin.Context) {
	id := c.Param("id")
	for _, l := range links {
		if fmt.Sprint(l.ID) == id {
			c.IndentedJSON(http.StatusOK, l)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "ID not found"})
}

func convertIDBase62(ID int) string {
	shortLink := string(base62.StdEncoding.Encode([]byte{byte(ID)}))
	return shortLink
}

func ConnectPostgres() (*sql.DB, error) {
	dsn := "host=localhost user=postgres password=postgres dbname=linksdb sslmode=disable port=5432"
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	return db, db.Ping()
}

func getTotalID(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, len(links))
}
