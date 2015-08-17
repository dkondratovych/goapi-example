package server

import (
	"os"
	"fmt"

	"github.com/seesawlabs/Dima-Kondravotych-Exercise/shared/config"
	"github.com/seesawlabs/Dima-Kondravotych-Exercise/shared/storages"
	"github.com/gin-gonic/gin"
	"github.com/seesawlabs/Dima-Kondravotych-Exercise/server/middleware"
)

type Server struct {
	Config *config.Config
	StorageProvider storages.IStorageProvider
	Router *gin.Engine
}

func NewServer(c *config.Config, sp storages.IStorageProvider) *Server {
	return &Server{
		Config: c,
		StorageProvider: sp,
	}
}

func(s *Server) Run() error {
	s.InitRouter()
	if err := s.InitHttpLogger(); err != nil {
		return err
	}
	s.SeDefaultMiddleware()
	s.SetRoutes()

	err := s.Router.Run(s.Config.Server.Port)
	if err != nil {
		return err
	}

	return nil
}

func(s *Server) InitRouter() {
	s.Router = gin.New()
}

func(s *Server) InitHttpLogger() error {
	file, err := os.OpenFile(s.Config.Server.HttpLogFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	if err != nil {
		return fmt.Errorf("Error openning http log file %v", err)
	}

	gin.DefaultWriter = file

	return nil
}

func(s *Server) SeDefaultMiddleware() error {
	s.Router.Use(gin.Logger())
	s.Router.Use(gin.Recovery())
	s.Router.Use(middleware.Cors())
	s.Router.Use(middleware.Application(s.Config))

	return nil
}

