package api

import (
	"main/db"
	"main/interceptor"
	"main/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func SetupTransactionAPI(router *gin.Engine) {
	transactionAPI := router.Group("api/v2")
	{
		transactionAPI.GET("/transaction", getTransaction)
		transactionAPI.POST("/transaction", interceptor.JwtVerify, createTransaction)
	}
}

type TransactionResult struct {
	ID uint
	Total float64
	Paid float64
	Change float64
	PaymentType string
	PaymentDetail string
	OrderList string
	Staff string
	CreatedAt time.Time
}

func getTransaction(c *gin.Context) {
	var result []TransactionResult
	db.GetDB().Model(&models.Transaction{}).Select("transactions.id, transactions.total, transactions.paid, transactions.change, transactions.payment_type, transactions.payment_detail, transactions.order_list, users.username as Staff, transactions.created_at").Joins("left join users on transactions.staff_id = users.id").Scan(&result)
	c.JSON(http.StatusOK, gin.H{"result": result})
}

func createTransaction(c *gin.Context) {
	var transaction models.Transaction
	if err := c.ShouldBind(&transaction); err == nil {
		transaction.StaffID = c.GetString("jwt_staff_id")
		transaction.CreatedAt = time.Now()
		db.GetDB().Create(&transaction)
		c.JSON(http.StatusOK, gin.H{"result": "ok", "transaction": transaction})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"result": "nok"})
	}
}
