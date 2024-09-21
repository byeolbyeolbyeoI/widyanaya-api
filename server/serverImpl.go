package server

import (
	"fmt"
	"github.com/byeolbyeolbyeoI/widyanaya-api/config"
	"github.com/byeolbyeolbyeoI/widyanaya-api/helper"
	userHandler "github.com/byeolbyeolbyeoI/widyanaya-api/internal/user/handler"
	userRepo "github.com/byeolbyeolbyeoI/widyanaya-api/internal/user/repository"
	userService "github.com/byeolbyeolbyeoI/widyanaya-api/internal/user/service"
	"github.com/gofiber/fiber/v2"
	supa "github.com/nedpals/supabase-go"
)

type Server struct {
	app    *fiber.App
	config *config.Config
	db     *supa.Client
	helper helper.HelperInstance
}

func NewServer(app *fiber.App, conf *config.Config, db *supa.Client) ServerInstance {
	h := helper.NewHelper()
	return &Server{app: app, config: conf, db: db, helper: h}
}

func (s *Server) Start() {
	userRepoInstance := userRepo.NewUserRepository(s.db, s.helper)
	userServiceInstance := userService.NewUserService(userRepoInstance, s.helper)
	userHandlerInstance := userHandler.NewUserHandler(userServiceInstance, s.helper)

	s.app = initializeRoutes(s.app, userHandlerInstance)
	if err := s.app.Listen(fmt.Sprintf(":%d", s.config.Server.Port)); err != nil {
		fmt.Println("error starting server:", err)
		return
	}
}
