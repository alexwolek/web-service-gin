package main

import (
	"example/web-service-gin/models"

	"net/http"

	"github.com/gin-gonic/gin"
)

// albums slice to seed record album data.
var albums = []models.Album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func main() {
	router := gin.Default()

	albums := router.Group("/albums")
	albums.GET("", getAlbums)

	router.Run("localhost:8080")
}
