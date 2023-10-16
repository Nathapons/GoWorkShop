package api

import "github.com/gin-gonic/gin"

func Setup(router *gin.Engine) {
	
	SetupAuthenAPI(router)
	SetupProductAPI(router)
	SetupTransactionAPI(router)
}