package server

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/facebook"

	"github.com/seesawlabs/Dima-Kondravotych-Exercise/server/middleware"
	"github.com/seesawlabs/Dima-Kondravotych-Exercise/shared/config"
	"github.com/seesawlabs/Dima-Kondravotych-Exercise/shared/storages"
)

type Server struct {
	Config          *config.Config
	StorageProvider storages.IStorageProvider
	Router          *gin.Engine
}

func NewServer(c *config.Config, sp storages.IStorageProvider) *Server {
	return &Server{
		Config:          c,
		StorageProvider: sp,
	}
}

func (s *Server) Run() error {
	s.InitRouter()
	if err := s.InitHttpLogger(); err != nil {
		return err
	}
	s.SeDefaultMiddleware()
	s.InitGothic()
	s.SetRoutes()

	err := s.Router.Run(s.Config.Server.Port)
	if err != nil {
		return err
	}

	return nil
}

func (s *Server) InitRouter() {
	s.Router = gin.New()
}

func (s *Server) InitGothic() {
	goth.UseProviders(
		facebook.New(s.Config.Facebook.AppId, s.Config.Facebook.Secret, s.Config.Facebook.CallbackURL),
	)
	// Since we are using only ONE provider, FACEBOOK - it doesn't have make sense to parse url for provider name
	gothic.GetProviderName = func(req *http.Request) (string, error) {
		return "facebook", nil
	}
}

func (s *Server) InitHttpLogger() error {
	file, err := os.OpenFile(s.Config.Server.HttpLogFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	if err != nil {
		return fmt.Errorf("Error openning http log file %v", err)
	}

	gin.DefaultWriter = file

	return nil
}

func (s *Server) SeDefaultMiddleware() error {
	s.Router.Use(gin.Logger())
	s.Router.Use(gin.Recovery())
	s.Router.Use(middleware.Cors())
	s.Router.Use(middleware.Application(s.Config))

	return nil
}
