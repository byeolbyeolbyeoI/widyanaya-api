package server

import (
	"fmt"
	"github.com/byeolbyeolbyeoI/icom/config"
	"github.com/gofiber/fiber/v2"
	"github.com/supabase-community/supabase-go"
)

type Server struct {
	app    *fiber.App
	config *config.Config
	db     *supabase.Client
}

func NewServer(app *fiber.App, conf *config.Config, db *supabase.Client) ServerInstance {
	return &Server{app: app, config: conf, db: db}
}

func (s *Server) Run() {
	s.app = initializeRoutes(s.app)
	if err := s.app.Listen(fmt.Sprintf(":%d", s.config.Server.Port)); err != nil {
		fmt.Println("error starting server:", err)
		return
	}
	fmt.Println("tes")
}
