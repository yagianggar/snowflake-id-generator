package main

import (
	"encoding/json"
	"log"
	"net/http"
	"snowflake-id-generator/utils/id"
	"strconv"
	"time"
)

type Response struct {
	SnowflakeID string    `json:"snowflake_id"`
	OrderID     uint64    `json:"order_id"`
	CreatedTime time.Time `json:"created_time"`
}

func main() {
	// HTTP handler to generate Snowflake SnowflakeID
	http.HandleFunc("/snowflake/generate", func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("Incoming request: method=%s, path=%s, remote=%s", r.Method, r.URL.Path, r.RemoteAddr)

		// Generate a Snowflake SnowflakeID
		generatedID, err := id.GetOrderNumber()
		if err != nil {
			log.Printf("Failed to generate order number: %v", err)
		}

		orderID, orderCreatedTime, err := id.ExtractIdInfo(generatedID)
		if err != nil {
			log.Printf("Failed to extract order number: %v", err)
		}

		// Create a response
		response := Response{
			SnowflakeID: generatedID,
			OrderID:     orderID,
			CreatedTime: orderCreatedTime,
		}

		// Set response headers
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		// Log the response data
		log.Printf("Response: id=%s, duration=%s", response.SnowflakeID, time.Since(start))

		// Send the JSON response
		if err := json.NewEncoder(w).Encode(response); err != nil {
			log.Printf("Failed to encode response: %v", err)
		}
	})

	// Start the HTTP server
	port := 8081
	log.Printf("Starting server on port %d...", port)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), nil))
}
