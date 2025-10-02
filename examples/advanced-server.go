
### 7. `advanced-server/advanced-server.go`
```go
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"
)

var (
	requestCount int
	mutex        sync.Mutex
	startTime    = time.Now()
)

// Middleware to log requests
func loggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		mutex.Lock()
		requestCount++
		count := requestCount
		mutex.Unlock()
		
		start := time.Now()
		log.Printf("📝 Request #%d: %s %s", count, r.Method, r.URL.Path)
		
		// Call the next handler
		next(w, r)
		
		duration := time.Since(start)
		log.Printf("✅ Request #%d completed in %v", count, duration)
	}
}

func main() {
	// Register routes with middleware
	http.HandleFunc("/", loggingMiddleware(rootHandler))
	http.HandleFunc("/api", loggingMiddleware(apiHandler))
	http.HandleFunc("/stats", loggingMiddleware(statsHandler))
	http.HandleFunc("/health", loggingMiddleware(healthHandler))
	http.HandleFunc("/greet", loggingMiddleware(greetHandler))

	// Server configuration
	port := ":8082"
	
	fmt.Println("🔥 Advanced Go Server")
	fmt.Println("====================")
	fmt.Printf("Server starting on http://localhost%s\n", port)
	fmt.Println("\nAvailable Endpoints:")
	fmt.Println("  GET /       - Welcome page")
	fmt.Println("  GET /api    - JSON API response")
	fmt.Println("  GET /stats  - Server statistics")
	fmt.Println("  GET /health - Health check")
	fmt.Println("  GET /greet  - Personalized greeting")
	fmt.Println("====================")
	fmt.Println("Features:")
	fmt.Println("  • Request logging middleware")
	fmt.Println("  • Request counting")
	fmt.Println("  • Performance timing")
	fmt.Println("  • Concurrent-safe operations")
	fmt.Println("====================")
	fmt.Println("Moringa AI Capstone - Advanced Server Example")

	// Start server
	log.Fatal(http.ListenAndServe(port, nil))
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `
	<!DOCTYPE html>
	<html>
	<head>
		<title>Advanced Go Server</title>
		<style>
			body { font-family: Arial, sans-serif; margin: 40px; }
			.container { max-width: 800px; margin: 0 auto; }
			.feature { background: #e8f4f8; padding: 15px; margin: 10px 0; border-radius: 5px; }
		</style>
	</head>
	<body>
		<div class="container">
			<h1>🔥 Advanced Go Server</h1>
			<p>This server demonstrates advanced features like middleware and request tracking.</p>
			
			<div class="feature">
				<strong>Features demonstrated:</strong>
				<ul>
					<li>Request logging middleware</li>
					<li>Request counting with mutex locks</li>
					<li>Performance timing</li>
					<li>Structured JSON responses</li>
					<li>Concurrent-safe operations</li>
				</ul>
			</div>
			
			<p>Check the <a href="/stats">/stats</a> endpoint for server statistics.</p>
			<p><em>Moringa AI Capstone - Advanced Example</em></p>
		</div>
	</body>
	</html>
	`)
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]interface{}{
		"message":   "Advanced Go Server API",
		"timestamp": time.Now(),
		"version":   "2.0.0",
		"features":  []string{"middleware", "logging", "metrics", "concurrency"},
	}
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func statsHandler(w http.ResponseWriter, r *http.Request) {
	stats := map[string]interface{}{
		"total_requests":   requestCount,
		"server_uptime":    time.Since(startTime).String(),
		"current_time":     time.Now().Format(time.RFC3339),
		"go_version":       "1.21+",
		"service":          "advanced-server",
		"project":          "Moringa AI Capstone",
	}
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(stats)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	health := map[string]string{
		"status":    "healthy",
		"timestamp": time.Now().Format(time.RFC3339),
		"service":   "advanced-server",
	}
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(health)
}

func greetHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "Friend"
	}
	
	response := map[string]string{
		"greeting":  fmt.Sprintf("Hello, %s!", name),
		"timestamp": time.Now().Format(time.RFC3339),
		"server":    "advanced-go-server",
	}
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}