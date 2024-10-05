package main

import (
	"fmt"
	"github.com/byeolbyeolbyeoI/widyanaya-api/config"
	"github.com/byeolbyeolbyeoI/widyanaya-api/database"
	"github.com/byeolbyeolbyeoI/widyanaya-api/server"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"strings"
)

func main() {
	app := fiber.New()
	conf := config.GetConfig()
	dbInstance := database.NewDatabase(conf)
	db := dbInstance.GetDatabase()

	app.Use(cors.New(cors.Config{
		// AllowOrigins:     "*",
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowCredentials: false,
		AllowMethods: strings.Join([]string{
			fiber.MethodGet,
			fiber.MethodPost,
			fiber.MethodHead,
			fiber.MethodPut,
			fiber.MethodDelete,
			fiber.MethodPatch,
		}, ","),
	}))

	serv := server.NewServer(app, conf, db)
	fmt.Println("listening on :8080")
	serv.Start()
}
