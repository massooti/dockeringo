package handler

import (
	"dockeringo/app"
	"net/http"
	"fmt"
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
	fmt.Printf("1")
	params := app.Parameters{"my-selenium2", runContainerRequest.Link, runContainerRequest.Guests}
	fmt.Printf("%+v\n", params)
	app.RunContainer4(params)
	
	c.JSON(200, gin.H{
		"message":   "short url created successfully",
		"short_url": "host" + "pa",
	})
}
