package main

import (
	"bank/config"
	"bank/routes"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	db := config.ConnectDB()
	defer db.Disconnect()
	routes.Route(r, db)

	r.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
