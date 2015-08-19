package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (o *OptionsHandler) AuthJwt(c *gin.Context) {
	c.Header("Allow", "POST")

	c.JSON(http.StatusCreated, TasksOptions{
		"POST": Method{
			Description: "Check user credentials and return jwt token",
			Parameters: Parameters{
				Parameter{
					"username": ParameterOption{
						"type":        "string",
						"description": "Username",
						"required":    true,
					},
					"password": ParameterOption{
						"type":        "string",
						"description": "Password",
						"required":    true,
					},
				},
			},
			Example: map[string]interface{}{
				"username": "Bender",
				"password": "molly",
			},
		},
	})
}
