package controllers

import (
	"example/web-service-gin/albums"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddAlbumsController(engine *gin.Engine) {
	albums := engine.Group("/albums")
	{
		albums.GET("", getAllAlbums)
		albums.GET("/:id", getAlbumByID)
		albums.POST("", addAlbum)
	}
}

// getAllAlbums responds with the list of all albums as JSON.
func getAllAlbums(c *gin.Context) {
	albums := albums.GetAllAlbums()

	c.IndentedJSON(http.StatusOK, albums)
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")
	album := albums.GetAlbumByID(id)
	if album == nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
	} else {
		c.IndentedJSON(http.StatusOK, album)
	}
}

// addAlbum adds an album from JSON received in the request body.
func addAlbum(c *gin.Context) {
	var newAlbum albums.Album

	// Call BindJSON to bind the received JSON to newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the collection.
	albums.AddNewAlbum(newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}
