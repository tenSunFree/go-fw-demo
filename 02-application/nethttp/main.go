package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	// Import shared models from the project
	"github.com/go-mizu/go-fw/pkg/models"
)

// Program execution entry point
func main() {
	// Create router
	mux := newRouter()
	// Create HTTP Server instance
	// Bind address and router handler together
	srv := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	// Output prompt message before starting
	fmt.Println("Server started successfully! Please visit: http://localhost:8080")
	// Start server
	// ListenAndServe will continue to block until the server encounters an error or is shut down.
	if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatal("Server failed to start: ", err)
	}
}

// newRouter is responsible for defining all routing rules and corresponding handler logic
func newRouter() http.Handler {
	// Create ServeMux router
	mux := http.NewServeMux()
	// Register root path route
	// Go 1.22+ supports direct use of "METHOD /path" pattern
	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		// Set response format to JSON
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		// Create response data
		response := models.ApiResponse{
			Code:    http.StatusOK,
			Message: "Query successful (Application pattern)",
			Data: models.UserData{
				ID:    2,
				Email: "app@example.com",
				Role:  "Admin, nethttp-02",
			},
		}
		// Encode struct to JSON and write back to Response
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, "failed to encode response", http.StatusInternalServerError)
			log.Println("failed to encode response:", err)
			return
		}
	})
	return mux
}
