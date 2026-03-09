package main

import (
	"fmt"
	"net/http"

	// Echo framework
	"github.com/labstack/echo/v4"
	// Shared models
	"github.com/go-mizu/go-fw/pkg/models"
)

func main() {
	// Initialize Echo instance
	// Similar to Gin's gin.Default(), echo.New() creates a new server instance
	e := echo.New()

	// Define routes
	// Note: Echo's Handler signature is func(echo.Context) error
	e.GET("/", func(c echo.Context) error {
		// Prepare response data
		response := models.ApiResponse{
			Code:    200,
			Message: "Query successful",
			Data: models.UserData{
				ID:    1,
				Email: "test@example.com",
				Role:  "Admin, echo", // Marked as echo version
			},
		}

		// Use c.JSON() for output
		// Unlike Gin, Echo requires you to return this result (because c.JSON returns an error)
		return c.JSON(http.StatusOK, response)
	})

	// Startup message
	fmt.Println("Server started successfully! Please visit: http://localhost:8080")

	// Start server
	// e.Start is equivalent to Gin's r.Run
	e.Logger.Fatal(e.Start(":8080"))
}
