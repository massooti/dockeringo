package handler

import (
	"dockeringo/app"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RunContainerRequest struct {
	Link   string `json:"link" binding:"required"`
	Guests int    `json:"guests" binding:"required"`
}

func ContainerInit(c *gin.Context) {
	/// Implementation to be added
	var runContainerRequest RunContainerRequest
	if err := c.ShouldBindJSON(&runContainerRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// fmt.Printf(runContainerRequest.Link)
	// pa := app.RunContainer2(runContainerRequest.Link)
	pa := app.ContainerRun()
	c.JSON(200, gin.H{
		"message":   "short url created successfully",
		"short_url": "host" + pa,
	})
}
