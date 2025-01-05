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
