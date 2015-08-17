package handlers

import (
	"time"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/seesawlabs/Dima-Kondravotych-Exercise/server/responses"
	"github.com/seesawlabs/Dima-Kondravotych-Exercise/shared/handlers"
	"github.com/dgrijalva/jwt-go"
)

type AuthHandler struct {
	handlers.CommonHandler
}

type AuthorizationRequest struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

const (
	FakeUserName = "Bender"
	FakePassword = "molly"
	FakeUserId = 11
)

func (ah *AuthHandler) Auth(c *gin.Context) {
	authRequest := &AuthorizationRequest{}
	if err := c.BindJSON(authRequest); err != nil {
		c.JSON(http.StatusInternalServerError, responses.ResponseError{
			ErrorCodeId: 54, // some fictional code
			DeveloperMessage: err.Error(),
			UserMessage: "An error occured whipe processing your request.",
		})
		return
	}

	if (authRequest.Password == FakePassword) || (authRequest.Password == FakePassword) {
		t := jwt.New(jwt.SigningMethodHS256)
		t.Claims["UserId"] = FakeUserId
		t.Claims["UserName"] = FakeUserName
		t.Claims["exp"] = time.Now().Add(time.Hour * 70).Unix()

		ts, err := t.SignedString(ah.Config.Server.JwtSecret)

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.ResponseError{
				ErrorCodeId: 21, // some fictional code
				DeveloperMessage: err.Error(),
				UserMessage: "An error occured while user authorization.",
			})
			return
		}

		c.JSON(http.StatusOK, map[string]string{"token": ts})

	} else {
		c.JSON(http.StatusUnauthorized, responses.ResponseError{
			ErrorCodeId: 22, // some fictional code
			DeveloperMessage: "Username or password is invalid.",
			UserMessage: "Username or password is invalid.",
		})
		return
	}
}


