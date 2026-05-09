package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	MsgHealthy = "Server is healthy"
)

func main() {
	router := gin.Default()

	router.Use(RequestID())
	router.Use(RequestLogger())

	router.GET("/health", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"status":  http.StatusText(http.StatusOK),
			"message": MsgHealthy,
		})
	})

	registerDashboardRoutes(router)
	registerDashboardRoutesInParallel(router)

	router.Run(":8080")
}
