package main

import (
	"dockeringo/handler"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hey Lets get started ...",
		})
	})

	r.POST("/start", func(c *gin.Context) {
		handler.Start(c)
	})

	r.POST("/stop", func(c *gin.Context) {
		handler.Stop(c)
	})

	err := r.Run(":9808")
	if err != nil {
		panic(fmt.Sprintf("Failed to start the web server - Error: %v", err))
	}
}
