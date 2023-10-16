package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupProductAPI(router *gin.Engine) {
	productAPI := router.Group("api/v2")
	{
		productAPI.GET("/product", getProduct)
		productAPI.POST("/product", createProduct)
	}
}

func getProduct(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"result": "products"})
}

func createProduct(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"result": "create product"})
}