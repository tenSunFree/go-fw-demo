package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	// Import the same common model as the net/http version.
	"github.com/go-mizu/go-fw/pkg/models"
	"github.com/go-mizu/mizu"
)

// Program execution entry point
func main() {
	fmt.Println("main")
	// Create App instance
	app := newApp()
	// Output startup prompt
	fmt.Println("Server started successfully! Please visit: http://localhost:8080")
	// Start server
	// If Listen returns an error, log it and exit
	if err := app.Listen(":8080"); err != nil {
		log.Fatal("Server failed to start: ", err)
	}
}

// newApp is responsible for creating a mizu App and registering all routes and handlers
func newApp() *mizu.App {
	fmt.Println("newApp")
	// Create mizu App
	app := mizu.New()
	// Register GET / route
	app.Get("/", func(c *mizu.Ctx) error {
		// Print request path to observe if request comes in
		fmt.Printf("Received request! Path is: %s\n", c.Request().URL.Path)
		// Create response data
		response := models.ApiResponse{
			Code:    http.StatusOK,
			Message: "Query successful (Application pattern)",
			Data: models.UserData{
				ID:    2,
				Email: "app@example.com",
				Role:  "Admin, mizu-02",
			},
		}
		// Output JSON to server terminal for development observation
		if err := json.NewEncoder(os.Stdout).Encode(response); err != nil {
			log.Println("failed to encode response to stdout:", err)
		}
		// Return JSON to client
		// c.JSON automatically sets Content-Type: application/json
		return c.JSON(http.StatusOK, response)
	})
	return app
}
