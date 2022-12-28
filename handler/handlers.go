package handler

import (
	"dockeringo/app"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type RunContainerRequest struct {
	Link   string `json:"link" binding:"required"`
	Guests int    `json:"guests" binding:"required"`
}

func Start(c *gin.Context) {
	/// Implementation to be added
	var runContainerRequest RunContainerRequest
	if err := c.ShouldBindJSON(&runContainerRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Printf("1")
	params := app.Parameters{"my-selenium", runContainerRequest.Link, runContainerRequest.Guests}
	app.StartAndRunContainers(params)

	c.JSON(200, gin.H{
		"message":   "short url created successfully",
		"short_url": "host" + "pa",
	})
}

func Stop(c *gin.Context) {
	// Stops and removes a container
	for i := 1; i <= 10; i++ {
		containerName := "cont-" + strconv.Itoa(i)
		app.StopAndRemoveContainer(containerName)
	}
	c.JSON(200, gin.H{
		"message": "All container is stopped... Bye Bye",
	})
}
