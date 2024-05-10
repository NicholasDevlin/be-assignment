package controller

import (
	"bank/dto/user"
	"bank/helper"
	errors "bank/helper/error"
	"bank/service"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService service.UserService
}

func NewUserController(userService service.UserService) *UserController {
	return &UserController{UserService: userService}
}

func (c *UserController) Create(ctx *gin.Context) {
	var user user.UserRequest
	if err := ctx.ShouldBindBodyWithJSON(&user); err != nil {
		helper.NewErrorResponse(ctx, errors.ERR_BAD_REQUEST)
		return
	}
	res, err := c.UserService.Create(ctx, user)
	if err != nil {
		helper.NewErrorResponse(ctx, err)
		return
	}
	helper.NewSuccessResponse(ctx, res)
}
