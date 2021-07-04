package main

import (
	"github.com/B-0-B-B-Y/alcoholve/server/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// POST /new - Return a new game object
	router.POST("/new", routes.NewGame)

	router.Run()
}
