package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/markbates/goth/gothic"

	"github.com/seesawlabs/Dima-Kondravotych-Exercise/server/responses"
)

func (ah *AuthHandler) FacebookAuth(c *gin.Context) {
	gothic.BeginAuthHandler(c.Writer, c.Request)
}

func (ah *AuthHandler) FacebookCallback(c *gin.Context) {
	user, err := gothic.CompleteUserAuth(c.Writer, c.Request)

	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.ResponseError{
			ErrorCodeId: 78, // some fictional code
			DeveloperMessage: err.Error(),
			UserMessage: "An error occured whipe processing your request.",
		})
		return
	}

	// @TODO Here we have got user access token and base information

	c.JSON(http.StatusOK, map[string]string{
		"UserAccessToken": user.AccessToken,
		"UserName": user.Name,
		"UserEmail": user.Email,
	})
}

