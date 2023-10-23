package api

import (
	"main/db"
	"main/models"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func SetupAuthenAPI(router *gin.Engine) {
	authenAPI := router.Group("api/v2")
	{
		authenAPI.POST("/login", login)
		authenAPI.POST("/register", register)
	}
}

func CheckPasswordHash(password string, hashPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
	return err == nil
}

func login(c *gin.Context) {
	var user models.User

	if c.ShouldBind(&user) == nil {
		var queryUser models.User
		err := db.GetDB().First(&queryUser, "username = ?", user.Username).Error
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"result": "ok", "error": err})
		} else if CheckPasswordHash(user.Password, queryUser.Password) == false {
			c.JSON(http.StatusUnauthorized, gin.H{"result": "unauthorized", "error": "invalid password"})
		} else {
			atClaims := jwt.MapClaims{}
			atClaims["id"] = queryUser.ID
			atClaims["username"] = queryUser.Username
			atClaims["level"] = queryUser.Level                                
			atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
			at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
			token, _ := at.SignedString([]byte("TokenSecret"))
			c.JSON(http.StatusOK, gin.H{"result": "ok", "token": token})
		}
	} else {
		c.JSON(http.StatusOK, gin.H{"result": "register"})
	}
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func register(c *gin.Context) {
	var user models.User
	if c.ShouldBind(&user) == nil {
		user.Password, _ = hashPassword(user.Password)
		user.CreatedAt = time.Now()
		if err := db.GetDB().Create(&user).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"result": "nok", "err": err})
		} else {
			c.JSON(http.StatusOK, gin.H{"result": "ok", "data": user})
		}

	} else {
		c.JSON(http.StatusOK, gin.H{"result": "register"})
	}
}
