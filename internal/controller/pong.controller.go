package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type PongController struct{}

func NewPongController() *PongController {
	return &PongController{}
}

func (uc *PongController) Pong(c *gin.Context) {
	name := c.Param("name")
	//name := c.DefaultQuery("name", "Guest") // Set default parameter name=Guest
	id := c.Query("uuid")

	c.JSON(http.StatusOK, gin.H{
		"message": "Pong",
		"name":    name,
		"id":      id,
	})
}
