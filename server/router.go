package server

import (
	"fmt"

	"github.com/seesawlabs/Dima-Kondravotych-Exercise/server/handlers"
	"github.com/seesawlabs/Dima-Kondravotych-Exercise/server/middleware"
	"github.com/seesawlabs/Dima-Kondravotych-Exercise/version"
)

func (s *Server) SetRoutes() {
	// Init Task handler
	taskHandler := &handlers.TaskHandler{}
	taskHandler.SetStorage(s.StorageProvider)

	opionsHandler := &handlers.OptionsHandler{}

	api := s.Router.Group(fmt.Sprintf("/api/%s", version.Version))
	api.Use(middleware.JwtAuth())
	api.POST("/tasks", taskHandler.Add)
	api.PUT("/tasks", taskHandler.Update)
	api.DELETE("/tasks", taskHandler.Delete)
	api.GET("/tasks/:id", taskHandler.Find)
	api.OPTIONS("/tasks", opionsHandler.Tasks)
	api.OPTIONS("/tasks/:id", opionsHandler.TaskById)

	// Init Auth handler
	authHandler := &handlers.AuthHandler{}
	authHandler.SetConfig(s.Config)
	auth := s.Router.Group("/auth")
	auth.POST("/jwt", authHandler.JwtAuth)
	auth.GET("/facebook", authHandler.FacebookAuth)
	auth.GET("/facebook/callback", authHandler.FacebookCallback)
	auth.OPTIONS("/jwt", opionsHandler.AuthJwt)
	auth.OPTIONS("/facebook", opionsHandler.AuthFacebook)
	auth.OPTIONS("/facebook/callback", opionsHandler.AuthFacebookCallback)
}
