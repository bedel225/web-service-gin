package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getAlbums(c *gin.Context) {
	c.JSON(http.StatusOK, albums)
}

func getAlbumByID(c *gin.Context) {
	for _, item := range albums {
		if item.ID == c.Param("id") {
			c.JSON(http.StatusOK, item)
		}
	}
}

func postAlbums(c *gin.Context) {
	var newAlbum album

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the slice.
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}
