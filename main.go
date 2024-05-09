package main

import (
	"bank/config"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	db := config.ConnectDB()
	defer db.Disconnect()

	r.Run(fmt.Sprintf(":%d", os.Getenv("PORT")))
}
