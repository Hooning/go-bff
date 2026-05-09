package main

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func RequestLogger() gin.HandlerFunc {
	return func(context *gin.Context) {
		start := time.Now()

		context.Next()

		requestID := context.GetString("request_id")

		log.Printf(
			"[REQUEST] (req:%s) %s %s | status: %d | duration: %v",
			requestID, context.Request.Method, context.Request.URL.Path, context.Writer.Status(), time.Since(start),
		)
	}
}

func RequestID() gin.HandlerFunc {
	return func(context *gin.Context) {
		requestID := uuid.New().String()

		context.Header("X-Request-ID", requestID)

		context.Set("request_id", requestID)

		context.Next()
	}
}
