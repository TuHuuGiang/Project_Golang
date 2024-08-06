package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/v1")
	{
		v1.GET("/ping/:name", Ping) // v1/ping/:name
		v1.POST("/ping", Ping)
		v1.PUT("/ping", Ping)
		v1.DELETE("/ping", Ping)
		v1.PATCH("/ping", Ping)
	}

	return r
}

func Ping(c *gin.Context) {
	name := c.Param("name")
	//name := c.DefaultQuery("name", "Guest") // Set default parameter name=Guest
	id := c.Query("uuid")

	c.JSON(http.StatusOK, gin.H{
		"message": "Pong",
		"name":    name,
		"id":      id,
	})
}
