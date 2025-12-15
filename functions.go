package main

import (
	"github/bedel225/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getAlbums(c *gin.Context) {
	albums, err := models.GetAllAlbum()

	if err != nil {
		panic("nothing to show")
	}
	c.JSON(http.StatusOK, albums)
}

func getAlbumByID(c *gin.Context) {
	for _, item := range models.Albums {
		if item.ID == c.Param("id") {
			c.JSON(http.StatusOK, item)
		}
	}
}

func postAlbums(c *gin.Context) {
	var newAlbum models.Album

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the slice.
	newAlbum.Save()
	c.IndentedJSON(http.StatusCreated, newAlbum)
}
