package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getAlbums(c *gin.Context) {
	c.JSON(http.StatusOK, albums)
}
