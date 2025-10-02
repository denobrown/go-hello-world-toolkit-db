package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

// Message represents a JSON response
type Message struct {
	Text      string    `json:"text"`
	Version   string    `json:"version"`
	Timestamp time.Time `json:"timestamp"`
	Service   string    `json:"service"`
}

// HealthCheck represents health status
type HealthCheck struct {
	Status    string `json:"status"`
	Timestamp string `json:"timestamp"`
	Uptime    string `json:"uptime"`
}

var startTime = time.Now()

func main() {
	// Register handlers
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/api", apiHandler)
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/info", infoHandler)
	http.HandleFunc("/time", timeHandler)

	// Server configuration
	port := ":8080"
	
	fmt.Println("🚀 Go HTTP Server Toolkit - Main Server")
	fmt.Println("=======================================")
	fmt.Printf("Server starting on http://localhost%s\n", port)
	fmt.Println("\nAvailable Endpoints:")
	fmt.Println("  GET /       - Welcome page")
	fmt.Println("  GET /api    - JSON API response")
	fmt.Println("  GET /health - Health check status")
	fmt.Println("  GET /info   - Server information")
	fmt.Println("  GET /time   - Current server time")
	fmt.Println("=======================================")
	fmt.Println("Moringa AI Capstone Project - Prompt Powered Learning")
	fmt.Println("Press Ctrl+C to stop the server")

	// Start server
	log.Fatal(http.ListenAndServe(port, nil))
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	html := `
	<!DOCTYPE html>
	<html>
	<head>
		<title>Go HTTP Server Toolkit</title>
		<style>
			body { font-family: Arial, sans-serif; margin: 40px; line-height: 1.6; }
			.container { max-width: 800px; margin: 0 auto; }
			.header { background: #00ADD8; color: white; padding: 20px; border-radius: 8px; }
			.endpoint { background: #f4f4f4; padding: 15px; margin: 10px 0; border-radius: 5px; }
			.code { background: #2d2d2d; color: #f8f8f2; padding: 15px; border-radius: 5px; }
		</style>
	</head>
	<body>
		<div class="container">
			<div class="header">
				<h1>🚀 Go HTTP Server Toolkit</h1>
				<p>Moringa AI Capstone Project - Prompt Powered Learning</p>
			</div>
			
			<h2>Welcome to the Main Server</h2>
			<p>This server demonstrates basic HTTP server capabilities in Go.</p>
			
			<h3>Available Endpoints:</h3>
			<div class="endpoint">
				<strong>GET /</strong> - This welcome page<br>
				<strong>GET /api</strong> - JSON API response<br>
				<strong>GET /health</strong> - Health check status<br>
				<strong>GET /info</strong> - Server information<br>
				<strong>GET /time</strong> - Current server time
			</div>
			
			<h3>Project Structure:</h3>
			<div class="code">
				go-http-server-toolkit/<br>
				├── basic-server/     # Simple HTTP server<br>
				├── advanced-server/  # Enhanced features<br>
				├── examples/         # Additional examples<br>
				├── docs/            # Documentation<br>
				└── main.go          # This server
			</div>
			
			<p><em>Explore the different server implementations in their respective directories!</em></p>
		</div>
	</body>
	</html>
	`
	fmt.Fprintf(w, html)
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	response := Message{
		Text:      "Welcome to the Go HTTP Server Toolkit API",
		Version:   "1.0.0",
		Timestamp: time.Now(),
		Service:   "main-server",
	}
	
	json.NewEncoder(w).Encode(response)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	health := HealthCheck{
		Status:    "healthy",
		Timestamp: time.Now().Format(time.RFC3339),
		Uptime:    time.Since(startTime).String(),
	}
	
	json.NewEncoder(w).Encode(health)
}

func infoHandler(w http.ResponseWriter, r *http.Request) {
	info := map[string]interface{}{
		"service":     "Go HTTP Server Toolkit - Main Server",
		"version":     "1.0.0",
		"go_version":  "1.21+",
		"project":     "Moringa AI Capstone",
		"start_time":  startTime.Format(time.RFC3339),
		"endpoints":   []string{"/", "/api", "/health", "/info", "/time"},
	}
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(info)
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now().Format("Monday, January 2, 2006 at 3:04:05 PM MST")
	fmt.Fprintf(w, "Server Time: %s\nUptime: %s", currentTime, time.Since(startTime))
}