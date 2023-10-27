package api

import (
	"main/db"
	"main/models"
	"net/http"
	"time"

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
	var transaction []models.Transaction
	db.GetDB().Find(&transaction)
	c.JSON(http.StatusOK, gin.H{"result": transaction})
}

func createTransaction(c *gin.Context) {
	var transaction models.Transaction
	if err := c.ShouldBind(&transaction); err == nil {
		transaction.CreatedAt = time.Now()
		db.GetDB().Create(&transaction)
		c.JSON(http.StatusOK, gin.H{"result": "ok", "transaction": transaction})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"result": "nok"})
	}
}
