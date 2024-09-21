package main

import (
	"fmt"
	"github.com/byeolbyeolbyeoI/widyanaya-api/config"
	"github.com/byeolbyeolbyeoI/widyanaya-api/database"
	"github.com/byeolbyeolbyeoI/widyanaya-api/server"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	conf := config.GetConfig()
	dbInstance := database.NewDatabase(conf)
	db := dbInstance.GetDatabase()

	serv := server.NewServer(app, conf, db)
	fmt.Println("listening on :8080")
	serv.Start()
}
