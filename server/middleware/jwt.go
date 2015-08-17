package middleware

import (
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/seesawlabs/Dima-Kondravotych-Exercise/server/responses"
)

func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		jwtSecretKey, ok := c.Get("JwtSecretKey")
		if !ok {
			c.JSON(http.StatusInternalServerError, responses.ResponseError{
				ErrorCodeId:      54, // some fictional code
				DeveloperMessage: "JWT secret token is missed.",
				UserMessage:      "Auth has failed.",
			})
			c.Abort()
		}

		var jsk interface{}
		jwtSecretKeyString := jwtSecretKey.(string)
		jwtSecretKeyByte := []byte(jwtSecretKeyString)
		jsk = jwtSecretKeyByte

		token, err := jwt.ParseFromRequest(c.Request, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Token signing method doesn't match")
			}

			return jsk, nil
		})

		if err == nil && token.Valid {
			c.Next()
		} else {
			c.JSON(http.StatusForbidden, responses.ResponseError{
				ErrorCodeId:      34, // some fictional code
				DeveloperMessage: err.Error(),
				UserMessage:      "Auth has failed.",
			})
			c.Abort()
		}
	}
}
