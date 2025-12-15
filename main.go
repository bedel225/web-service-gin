package main

import (
	"github/bedel225/db"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()

	router := gin.Default()

	router.GET("/albums", getAlbums)
	//router.GET("/albums/id/:id", getAlbumByID)
	router.POST("/albums", postAlbums)

	router.Run("localhost:8000")
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
