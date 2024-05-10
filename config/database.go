package config

import (
	"bank/prisma/db"

	"github.com/labstack/gommon/log"
)

func ConnectDB() (*db.PrismaClient) {
	client := db.NewClient()
	if err := client.Prisma.Connect(); err != nil {
		return nil
	}
	log.Info("Connected to database")
	return client
}
