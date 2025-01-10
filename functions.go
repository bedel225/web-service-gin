package main

import (
	"fmt"
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

func postAlbums(c *gin.Context) {

	var newAlbum models.Album

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}
	fmt.Println(newAlbum)
	newAlbum.Save()

	c.IndentedJSON(http.StatusCreated, newAlbum)
}
