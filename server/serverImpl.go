package server

import (
	"fmt"
	"github.com/byeolbyeolbyeoI/icom/config"
	"github.com/byeolbyeolbyeoI/icom/database"
	"github.com/gofiber/fiber/v2"
)

type Server struct {
	app    *fiber.App
	config *config.Config
	db     *database.DatabaseInstance
}

func NewServer(app *fiber.App, conf *config.Config, db *database.DatabaseInstance) ServerInstance {
	return &Server{app: app, config: conf, db: db}
}

func (s *Server) Run() {
	if err := s.app.Listen(fmt.Sprintf(":%s", s.config.Server.Port)); err != nil {
		fmt.Println("error starting server:", err)
		return
	}
}
