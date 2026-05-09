package main

import (
	"log"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

type Result[T any] struct {
	data T
	err  error
}

func registerDashboardRoutes(router *gin.Engine) {
	router.GET("/api/dashboard/overview", func(context *gin.Context) {
		user, _ := fetchUser()
		metrics, _ := fetchMetrics()
		alerts, _ := fetchAlerts()

		context.JSON(http.StatusOK, gin.H{
			"user":    user,
			"metrics": metrics,
			"alerts":  alerts,
		})
	})
}

func registerDashboardRoutesInParallel(router *gin.Engine) {
	router.GET("/api/dashboard/overview-parallel", func(context *gin.Context) {
		requestID := context.GetString("request_id")

		var waitGroup sync.WaitGroup

		userChannel := make(chan Result[User], 1)
		metricsChannel := make(chan Result[[]Metric], 1)
		alertsChannel := make(chan Result[[]Alert], 1)

		// Variable to store results from each goroutine
		// var (
		// 	user      User
		// 	metrics   []Metric
		// 	alerts    []Alert
		// 	waitGroup sync.WaitGroup
		// )

		// waitGroup tracks when all goroutines have completed
		// Similar as Promise.all() in JavaScript
		waitGroup.Add(3)

		go func() {
			defer waitGroup.Done()
			data, err := fetchUser()
			userChannel <- Result[User]{data, err}
		}()

		go func() {
			defer waitGroup.Done()
			data, err := fetchMetrics()
			metricsChannel <- Result[[]Metric]{data, err}
		}()

		go func() {
			defer waitGroup.Done()
			data, err := fetchAlerts()
			alertsChannel <- Result[[]Alert]{data, err}
		}()

		// Block until all 3 goroutines finish their work
		waitGroup.Wait()

		userResult := <-userChannel
		metricsResult := <-metricsChannel
		alertsResult := <-alertsChannel

		response := gin.H{}

		if userResult.err != nil {
			log.Printf("[ERROR] (req:%s) user service failed: %v", requestID, userResult.err)
			response["user"] = nil
		} else {
			response["user"] = userResult.data
		}

		if metricsResult.err != nil {
			log.Printf("[ERROR] (req:%s) metrics service failed: %v", requestID, metricsResult.err)
			response["metrics"] = nil
		} else {
			response["metrics"] = metricsResult.data
		}

		if alertsResult.err != nil {
			log.Printf("[ERROR] (req:%s) alerts service failed: %v", requestID, alertsResult.err)
			response["alerts"] = nil
		} else {
			response["alerts"] = alertsResult.data
		}

		context.JSON(http.StatusOK, response)

		// context.JSON(http.StatusOK, gin.H{
		// 	"user":    user,
		// 	"metrics": metrics,
		// 	"alerts":  alerts,
		// })
	})
}
