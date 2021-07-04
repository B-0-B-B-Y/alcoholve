package main

import (
	"github.com/B-0-B-B-Y/alcoholve/server/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// POST /new - Return a new game object
	router.POST("/new", routes.NewGame)

	// POST /join - Allow a client to join a game
	router.POST("/join", routes.Join)

	// GET /ping - Default fallback route to verify server is up
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.Run()
}
