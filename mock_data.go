package main

import (
	"fmt"
	"time"
)

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Role string `json:"role"`
}

type Metric struct {
	Label string `json:"label"`
	Value int    `json:"value"`
}

type Alert struct {
	Level   string `json:"level"`
	Message string `json:"message"`
}

func fetchUser() (User, error) {
	time.Sleep(200 * time.Millisecond)  // simulate DB call
	return User{ID: "123", Name: "Hoon Cho", Role: "Software Engineer"}, nil
}

func fetchMetrics() ([]Metric, error) {
	time.Sleep(150 * time.Millisecond)  // simulate slower query
	
	// Successfully fetched metrics
	// return []Metric{
	// 	{Label: "CPU Usage", Value: 75},
	// 	{Label: "Memory Usage", Value: 60},
	// }

	// Simulate a failure
	return nil, fmt.Errorf("Metrics service is currently unavailable")
}

func fetchAlerts() ([]Alert, error) {
	time.Sleep(100 * time.Millisecond)  // simulate DB call
	return []Alert{
		{Level: "Warning", Message: "High CPU usage detected"},
		{Level: "Info", Message: "Memory usage is within normal limits"},
	}, nil
}
