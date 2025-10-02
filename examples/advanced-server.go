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
)

// Middleware to log requests
func loggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		mutex.Lock()
		requestCount++
		count := requestCount
		mutex.Unlock()
		
		start := time.Now()
		log.Printf("Request #%d: %s %s", count, r.Method, r.URL.Path)
		
		next(w, r)
		
		duration := time.Since(start)
		log.Printf("Request #%d completed in %v", count, duration)
	}
}

// Stats endpoint
func statsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	stats := map[string]interface{}{
		"total_requests": requestCount,
		"server_uptime":  time.Since(startTime).String(),
		"timestamp":      time.Now(),
	}
	
	json.NewEncoder(w).Encode(stats)
}

// Dynamic greeting endpoint
func greetHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "Friend"
	}
	
	response := map[string]string{
		"greeting": fmt.Sprintf("Hello, %s!", name),
		"language": "Go",
	}
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

var startTime = time.Now()

func main() {
	// Register routes with middleware
	http.HandleFunc("/", loggingMiddleware(helloHandler))
	http.HandleFunc("/api", loggingMiddleware(apiHandler))
	http.HandleFunc("/health", loggingMiddleware(healthHandler))
	http.HandleFunc("/time", loggingMiddleware(timeHandler))
	http.HandleFunc("/stats", loggingMiddleware(statsHandler))
	http.HandleFunc("/greet", loggingMiddleware(greetHandler))
	
	port := ":8080"
	
	fmt.Printf("🚀 Advanced Go Server starting on http://localhost%s\n", port)
	fmt.Println("==========================================")
	fmt.Println("Available endpoints:")
	fmt.Println("  GET /              - Welcome page")
	fmt.Println("  GET /api           - JSON API")
	fmt.Println("  GET /health        - Health check")
	fmt.Println("  GET /time          - Current time")
	fmt.Println("  GET /stats         - Server statistics")
	fmt.Println("  GET /greet?name=John - Personalized greeting")
	fmt.Println("==========================================")
	
	log.Fatal(http.ListenAndServe(port, nil))
}