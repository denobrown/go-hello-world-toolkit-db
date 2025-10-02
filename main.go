package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

// Message struct for JSON responses
type Message struct {
	Text      string    `json:"text"`
	Version   string    `json:"version"`
	Timestamp time.Time `json:"timestamp"`
}

// HealthCheck struct for health endpoint
type HealthCheck struct {
	Status    string `json:"status"`
	Timestamp string `json:"timestamp"`
}

// Handler for the root route
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `
	<!DOCTYPE html>
	<html>
	<head>
		<title>Go Hello World</title>
		<style>
			body { font-family: Arial, sans-serif; margin: 40px; }
			.container { max-width: 800px; margin: 0 auto; }
			.code { background: #f4f4f4; padding: 10px; border-radius: 5px; }
		</style>
	</head>
	<body>
		<div class="container">
			<h1>🚀 Welcome to Go!</h1>
			<p>This is a simple HTTP server written in Go.</p>
			<div class="code">
				<strong>Endpoints:</strong><br>
				• <a href="/">/</a> - This page<br>
				• <a href="/api">/api</a> - JSON API<br>
				• <a href="/health">/health</a> - Health check<br>
				• <a href="/time">/time</a> - Current time
			</div>
			<p><em>Server is running on Go version 1.21+</em></p>
		</div>
	</body>
	</html>
	`)
}

// Handler for API endpoint
func apiHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	response := Message{
		Text:      "Hello from the Go API!",
		Version:   "1.0",
		Timestamp: time.Now(),
	}
	
	json.NewEncoder(w).Encode(response)
}

// Handler for health check
func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	health := HealthCheck{
		Status:    "healthy",
		Timestamp: time.Now().Format(time.RFC3339),
	}
	
	json.NewEncoder(w).Encode(health)
}

// Handler for current time
func timeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Current server time: %s", time.Now().Format("2006-01-02 15:04:05 MST"))
}

func main() {
	// Define routes
	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/api", apiHandler)
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/time", timeHandler)
	
	// Server configuration
	port := ":8080"
	
	// Print startup message
	fmt.Printf("Starting Go HTTP server on http://localhost%s\n", port)
	fmt.Println("==========================================")
	fmt.Println("Available endpoints:")
	fmt.Println("  GET /      - Welcome page")
	fmt.Println("  GET /api   - JSON API response")
	fmt.Println("  GET /health- Health check status")
	fmt.Println("  GET /time  - Current server time")
	fmt.Println("==========================================")
	fmt.Println("Press Ctrl+C to stop the server")
	
	// Start server
	log.Fatal(http.ListenAndServe(port, nil))
}