package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/seesawlabs/Dima-Kondravotych-Exercise/shared/config"
)

func Application(conf *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		// @TODO just set secret key here, later while application/appconfig can be added
		c.Set("JwtSecretKey", conf.Server.JwtSecret)
		c.Next()
	}
}