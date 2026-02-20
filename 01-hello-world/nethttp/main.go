// Define package name
// Tell the Go compiler that this is an executable program entry point, not just a library for import.
package main

import (
	// Format I/O (for printing text)
	"fmt"
	// Log (for logging errors)
	"log"
	// Go's official core network protocol library
	// Network HTTP (for creating web servers)
	// net/http is the foundation of all Go frameworks. Whether it's Gin or Fiber, they all ultimately call the logic here.
	"net/http"
	// Encoding JSON (for JSON serialization/deserialization)
	"encoding/json"
)

// Define standard response structure (usually placed globally or in a dedicated folder)
type ApiResponse struct {
	Code    int         `json:"code"`    // Custom status code
	Message string      `json:"message"` // Message for frontend display
	Data    interface{} `json:"data"`    // Actual data content, can be anything
}

// Define specific business data structure
type UserData struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
	Role  string `json:"role"`
}

// Program execution entry point
func main() {
	// Create router
	// Create a new ServeMux (a multiplexer for routing HTTP requests)
	// It's like a receptionist at a restaurant entrance, responsible for bringing incoming guests (Requests) to the correct table (Handler).
	mux := http.NewServeMux()

	// Set path and behavior
	// Register a handler function for the root path "/"
	// GET /: This is new syntax after Go 1.22, directly specifying "method" and "path".
	// w: ResponseWriter - used to send responses back to the client
	// r: Request - contains all information about the incoming HTTP request
	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		// Write response content
		// Send the string through w back to the browser. Fprintln's F stands for File (in Linux philosophy, network connections are also file streams).
		// fmt.Fprintln(w, "hello, world!")

		// Set Header, tell the browser this is JSON format
		w.Header().Set("Content-Type", "application/json")

		// Define the data structure you want to return (usually use struct)
		// data := map[string]interface{}{
		// 	"status":  "success",
		// 	"message": "hello, world! This is JSON format",
		// 	"data": map[string]int{
		// 		"id": 123,
		// 	},
		// }
		// Use json package to convert map or struct to JSON and write to response
		// json.NewEncoder(w).Encode(data)

		// Instantiate struct
		response := ApiResponse{
			Code:    200,
			Message: "Query successful",
			Data: UserData{
				ID:    1,
				Email: "test@example.com",
				Role:  "Admin",
			},
		}
		// Encode and return
		json.NewEncoder(w).Encode(response)
	})

	// Print a prompt before starting to let yourself know where to click the URL
	fmt.Println("Server started successfully! Please visit: http://localhost:8080")

	// Start and listen
	// Listen on port 8080 and handle requests using the mux router
	// This line will block here, starting to guard port 8080 on the computer. If someone knocks on the door, the mux receptionist will be called to handle it.
	// This line will block continuously until an error occurs
	// If an error occurs, log.Fatal will print the timestamp and error, then terminate the program directly
	log.Fatal(http.ListenAndServe(":8080", mux))
	// http.ListenAndServe(":8080", mux)
}
