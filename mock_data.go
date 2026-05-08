package main

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

func fetchUser() User {
	return User{ID: "123", Name: "Hoon Cho", Role: "Software Engineer"}
}

func fetchMetrics() []Metric {
	return []Metric{
		{Label: "CPU Usage", Value: 75},
		{Label: "Memory Usage", Value: 60},
	}
}

func fetchAlerts() []Alert {
	return []Alert{
		{Level: "Warning", Message: "High CPU usage detected"},
		{Level: "Info", Message: "Memory usage is within normal limits"},
	}
}
