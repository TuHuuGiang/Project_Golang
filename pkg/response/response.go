package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ResponseData struct {
	Code    int         `json:"code"`    // status code
	Message string      `json:"message"` // Thông báo lỗi
	Data    interface{} `json:"data"`    // Dữ liệu trả về
}

func SuccessResponse(c *gin.Context, code int, data interface{}) {
	c.JSON(http.StatusOK, ResponseData{
		Code:    code,
		Message: msg[code],
		Data:    data,
	})
}

func ErrorResponse(c *gin.Context, code int, message string) {
	c.JSON(http.StatusInternalServerError, ResponseData{
		Code:    code,
		Message: msg[code],
		Data:    nil,
	})
}
