package helper

import (
	"bank/helper/errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BaseResponse struct {
	Success  bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewSuccessResponse(c *gin.Context, data interface{}) {
	c.IndentedJSON(http.StatusOK, BaseResponse{
		Success:  true,
		Message: "Success",
		Data:    data,
	})
}

func NewErrorResponse(c *gin.Context, err error) {
	c.IndentedJSON(errors.GetErrorCode(err), BaseResponse{
		Success: false,
		Message: err.Error(),
		Data:    nil,
	})
}
