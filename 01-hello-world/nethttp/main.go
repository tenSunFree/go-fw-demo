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
)

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
		fmt.Fprintln(w, "hello, world!")
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
