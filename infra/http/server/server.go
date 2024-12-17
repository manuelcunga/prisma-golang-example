package server

import (
	"errors"
	"log"
	"prisma-golang/infra/database"
	"prisma-golang/infra/http/router"

	"github.com/gin-gonic/gin"
)

type Server struct {
	gin *gin.Engine
}

func NewServer(g *gin.Engine) *Server {
	return &Server{gin: g}
}

func (s *Server) StartServer() error {
	gin.SetMode(gin.DebugMode)

	database, err := database.ConnectDB()
	if err != nil {
		return errors.New("error to connect database")
	}

	routes := router.NewRouter(s.gin)
	routes.RouterSetup(*database)

	if err := s.gin.Run(":5000"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

	return nil
}
