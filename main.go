package main

import (
	"log"

	"github.com/Devisrisamidurai/go-rest-api/src/albums"
	"github.com/gin-gonic/gin"
)

func main() {

	// initialise Gin router:
	router := gin.Default()
	router.GET("/albums", albums.GetAlbums)
	router.POST("/albums", albums.PostAlbums)

	// run the server on port 3000:
	err := router.Run(":5000")
	if err != nil {
		log.Fatalf("[Error] failed to start Gin server due to: " + err.Error())
	}

}
