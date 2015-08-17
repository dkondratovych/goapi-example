package server

import (
	"fmt"

	"github.com/seesawlabs/Dima-Kondravotych-Exercise/version"
	"github.com/seesawlabs/Dima-Kondravotych-Exercise/server/handlers"
	"github.com/seesawlabs/Dima-Kondravotych-Exercise/server/middleware"
)

func(s *Server) SetRoutes() {
	// Init Task handler
	taskHandler := (&handlers.TaskHandler{})
	taskHandler.SetStorage(s.StorageProvider)

	api := s.Router.Group(fmt.Sprintf("/api/%s", version.Version))
	api.Use(middleware.JwtAuth())
	api.POST("/tasks", taskHandler.Add)
	api.GET("/tasks/:id", taskHandler.Find)
	api.PUT("/tasks", taskHandler.Update)
	api.DELETE("/tasks", taskHandler.Delete)

	// Init Auth handler
	authHandler := (&handlers.AuthHandler{})
	authHandler.SetConfig(s.Config)
	s.Router.POST("/auth/jwt", authHandler.JwtAuth)

	s.Router.GET("/auth/facebook", authHandler.FacebookAuth)
	s.Router.GET("/auth/facebook/callback", authHandler.FacebookCallback)
}