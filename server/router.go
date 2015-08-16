package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/seesawlabs/Dima-Kondravotych-Exercise/version"
	"github.com/seesawlabs/Dima-Kondravotych-Exercise/server/handlers"
)

func(s *Server) InitRouter() {
	s.Router = gin.Default()

	taskHandler := (&handlers.TaskHandler{})
	taskHandler.SetStorage(s.StorageProvider)

	v := s.Router.Group(fmt.Sprintf("/api/%s", version.Version))

	v.POST("/tasks", taskHandler.Add)
	v.GET("/tasks/:id", taskHandler.Find)
	v.PUT("/tasks", taskHandler.Update)
	v.DELETE("/tasks", taskHandler.Delete)

}