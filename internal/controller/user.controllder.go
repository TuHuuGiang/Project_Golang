package controller

import (
	"github.com/gin-gonic/gin"
	"go-ecommerce-be-api/internal/service"
	"go-ecommerce-be-api/pkg/response"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController() *UserController {
	return &UserController{
		userService: service.NewUserService(),
	}
}

func (uc *UserController) GetUserById(c *gin.Context) {
	response.SuccessResponse(c, response.CodeSuccess, []string{"cr7", "thg"})

	response.ErrorResponse(c, response.ErrorCodeParamInvalid, "Email is invalid")
}
