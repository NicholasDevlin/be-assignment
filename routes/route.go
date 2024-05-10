package routes

import (
	"bank/prisma/db"

	"github.com/gin-gonic/gin"
)

func Route(c *gin.Engine, db *db.PrismaClient) {
	UserRoute(c, db)
}
