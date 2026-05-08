package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/health", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"status":  "Ok",
			"message": "Server is healthy",
		})
	})

	registerDashboardRoutes(router)

	router.Run(":8080")
}
