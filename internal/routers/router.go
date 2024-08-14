package routers

import (
	"github.com/gin-gonic/gin"
	"go-ecommerce-be-api/internal/controller"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/v1")
	{
		v1.GET("/ping", controller.NewPongController().Pong)
		v1.GET("user/1", controller.NewUserController().GetUserById)
		//v1.GET("/ping/:name", controller.Pong) // v1/ping/:name
		//v1.POST("/ping", Pong)
		//v1.PUT("/ping", Pong)
		//v1.DELETE("/ping", Pong)
		//v1.PATCH("/ping", Pong)
	}

	return r
}

//func Pong(c *gin.Context) {
//	name := c.Param("name")
//	//name := c.DefaultQuery("name", "Guest") // Set default parameter name=Guest
//	id := c.Query("uuid")
//
//	c.JSON(http.StatusOK, gin.H{
//		"message": "Pong",
//		"name":    name,
//		"id":      id,
//	})
//}
