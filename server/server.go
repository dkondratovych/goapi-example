package server

import (
	"github.com/seesawlabs/Dima-Kondravotych-Exercise/shared/config"
	"github.com/seesawlabs/Dima-Kondravotych-Exercise/shared/storages"
	"github.com/gin-gonic/gin"
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

	err := s.Router.Run(s.Config.Server.Port)
	if err != nil {
		return err
	}

	return nil
}

