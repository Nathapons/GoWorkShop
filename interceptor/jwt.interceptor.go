package interceptor

import (
	"fmt"
	"main/models"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var secretKey string = "qwerty1234"

func JwtVerify(c *gin.Context) {
	tokenString := strings.Split(c.Request.Header["Authorization"][0], " ")[1]
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method : %v", token.Header["alg"])
		}
		return []byte(secretKey), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Change data type to string
		staffID := fmt.Sprintf("%v", claims["id"])
		username := fmt.Sprintf("%v", claims["username"])
		level := fmt.Sprintf("%v", claims["level"])

		// Set data to gin.Context
		c.Set("jwt_staff_id", staffID)
		c.Set("jwt_username", username)
		c.Set("level", level)

		// Next to GIN Router
		c.Next()
	} else {
		c.JSON(http.StatusOK, gin.H{"result": "nok", "error": err})
		c.Abort()
	}
}

func JwtSign(payload models.User) string {
	atClaims := jwt.MapClaims{}
	atClaims["id"] = payload.ID
	atClaims["username"] = payload.Username
	atClaims["level"] = payload.Level
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, _ := at.SignedString([]byte(secretKey))
	return token
}
