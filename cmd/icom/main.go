package main

import (
	"github.com/byeolbyeolbyeoI/icom/config"
	"github.com/byeolbyeolbyeoI/icom/database"
	"github.com/byeolbyeolbyeoI/icom/server"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	conf := config.GetConfig()
	dbInstance := database.NewDatabase(conf)
	db := dbInstance.GetDatabase()

	server.NewServer(app, conf, db)
}