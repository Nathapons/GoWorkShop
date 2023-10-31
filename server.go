package main

import (
	"fmt"
	"main/apis"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
)

func main() {
	port := os.Getenv("PORT")
	router := gin.Default()
	// Set up CORS middleware options
	config := cors.Config{
		Origins:        "*",
		RequestHeaders: "Origin, Authorization, Content-Type",

		Methods:         "GET, POST, PUT, DELETE",
		Credentials:     false,
		ValidateHeaders: false,
		MaxAge:          1 * time.Minute,
	}

	router.Use(cors.Middleware(config))
	router.Static("/images", "./uploaded/images")

	api.Setup(router)
	fmt.Printf("Running port %v \n", port)
	router.Run(":" + port)
}
