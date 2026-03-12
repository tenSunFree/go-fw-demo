package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	// Shared models
	"github.com/go-mizu/go-fw/pkg/models"
)

func main() {
	// Initialize Chi router
	// Similar to Echo's echo.New(), chi.NewRouter() creates a new router instance
	r := chi.NewRouter()

	// Define routes
	// Note: Chi's Handler signature is func(http.ResponseWriter, *http.Request)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		// Prepare response data
		response := models.ApiResponse{
			Code:    200,
			Message: "Query successful",
			Data: models.UserData{
				ID:    1,
				Email: "test@example.com",
				Role:  "Admin, chi", // Marked as chi version
			},
		}

		// Set response header to JSON
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		// Output JSON response
		// Unlike Echo's c.JSON(), Chi needs manual JSON encoding
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
			return
		}
	})

	// Startup message
	fmt.Println("Server started successfully! Please visit: http://localhost:8080")

	// Start server
	// http.ListenAndServe is equivalent to Echo's e.Start
	if err := http.ListenAndServe(":8080", r); err != nil {
		panic(err)
	}
}
