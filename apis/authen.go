package api

import (
	"main/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupAuthenAPI(router *gin.Engine) {
	authenAPI := router.Group("api/v2")
	{
		authenAPI.POST("/login", login)
		authenAPI.POST("/register", register)
	}
}

func login(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"result": "Login"})
}

func register(c *gin.Context) {
	var user models.User
	if c.ShouldBind(&user) == nil {
		c.JSON(http.StatusOK, gin.H{"result": "register", "data": user})
	}
	c.JSON(http.StatusOK, gin.H{"result": "register"})
}
