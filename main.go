package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	api := router.Group("/api")
	{
		api.GET("/get/:id", func(c *gin.Context) {
			id := c.Param("id")
			c.JSON(http.StatusOK, gin.H{
				"message": "ping pong successful: " + id,
			})
		})

		api.POST("/create", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H {
				"message": "Create new post sucessful",
			})
		})
	}

	router.Run(":8080")
}
