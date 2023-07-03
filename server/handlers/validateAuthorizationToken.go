package handlers

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"jwt-filter/server/conf"
	"log"
	"net/http"
	"strings"
)

func ValidateAuthorizationToken(c *gin.Context) {
	var statusCode = http.StatusUnauthorized
	var msg = ""

	AuthorizationHeader := c.GetHeader("Authorization")
	if !strings.HasPrefix(AuthorizationHeader, "Bearer ") {
		c.JSON(statusCode, gin.H{"error": msg})
		return
	}

	tokenString := strings.Split(AuthorizationHeader, " ")[1]
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(conf.JwtSecretKey), nil
	})

	var userId string
	if token.Valid {
		statusCode = http.StatusOK
		msg = "Good token"

		tokenClaims := token.Claims.(jwt.MapClaims)
		var iss float64 = tokenClaims["iss"].(float64)

		userId = fmt.Sprintf("%.0f", iss)
		if userId == "" {
			msg = "Good token [WARNING!!! No `iss` field]"
			log.Println("[WARNING!!! No `iss` field in token found]")
		}
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			msg = "That's not even a token"
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			msg = "Token expired"
		} else {
			msg = "Couldn't handle this token:"
		}
	} else {
		msg = "Couldn't handle this token:"
	}

	c.Header(conf.UserIdHeader, strings.TrimSpace(userId))
	c.JSON(statusCode, gin.H{"msg": msg})
}
