package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Register your routes and handlers here
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})

	r.Run(":5000") // Run the HTTP server on port 5000
}
