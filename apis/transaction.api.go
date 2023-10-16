package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupTransactionAPI(router *gin.Engine) {
	transactionAPI := router.Group("api/v2")
	{
		transactionAPI.GET("/transaction", getTransaction)
		transactionAPI.POST("/transaction", createTransaction)
	}
}

func getTransaction(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"result": "Transaction"})
}

func createTransaction(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"result": "create transaction"})
}
