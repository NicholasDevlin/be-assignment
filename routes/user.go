package routes

import (
	"bank/controller"
	"bank/prisma/db"
	repository "bank/repository/implement"
	service "bank/service/implement"

	"github.com/gin-gonic/gin"
)

func UserRoute(c *gin.Engine, db *db.PrismaClient) {
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(userService)

	c.POST("/user", userController.Create)
	c.POST("/user/login", userController.Login)
}
