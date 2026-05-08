package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func registerDashboardRoutes(router *gin.Engine) {
	router.GET("/api/dashboard/overview", func(context *gin.Context) {
		user := fetchUser()
		metrics := fetchMetrics()
		alerts := fetchAlerts()

		context.JSON(http.StatusOK, gin.H{
			"user":    user,
			"metrics": metrics,
			"alerts":  alerts,
		})
	})
}
