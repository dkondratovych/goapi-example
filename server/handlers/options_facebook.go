package handlers

import (
	"github.com/gin-gonic/gin"
)

func (o *OptionsHandler) AuthFacebook(c *gin.Context) {
	c.Header("Allow", "GET")
}

func (o *OptionsHandler) AuthFacebookCallback(c *gin.Context) {
	c.Header("Allow", "GET")
}
