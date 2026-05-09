package main

import (
	"net/http"
	"sync"

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

func registerDashboardRoutesInParallel(router *gin.Engine) {
	router.GET("/api/dashboard/overview-parallel", func(context *gin.Context) {
		// Variable to store results from each goroutine
		var (
			user      User
			metrics   []Metric
			alerts    []Alert
			waitGroup sync.WaitGroup
		)

		// waitGroup tracks when all goroutines have completed
		// Similar as Promise.all() in JavaScript
		waitGroup.Add(3)

		go func() {
			defer waitGroup.Done()
			user = fetchUser()
		}()

		go func() {
			defer waitGroup.Done()
			metrics = fetchMetrics()
		}()

		go func() {
			defer waitGroup.Done()
			alerts = fetchAlerts()
		}()

		// Block until all 3 goroutines finish their work
		waitGroup.Wait()

		context.JSON(http.StatusOK, gin.H{
			"user":    user,
			"metrics": metrics,
			"alerts":  alerts,
		})
	})
}
