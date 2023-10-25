package api

import (
	"net/http"
	"main/interceptor"

	"github.com/gin-gonic/gin"
)

// SetupProductAPI - call this method to setup product route group
func SetupProductAPI(router *gin.Engine) {
	productAPI := router.Group("api/v2")
	{
		productAPI.GET("/product", interceptor.JwtVerify, getProduct)
		productAPI.POST("/product", createProduct)
	}
}

func getProduct(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"result": "products", "username": c.GetString("jwt_username"), "level": c.GetString("level")})
}

func createProduct(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"result": "create product"})
}
