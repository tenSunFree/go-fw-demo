package main

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	// Shared models
	"github.com/go-mizu/go-fw/pkg/models"
)

func main() {
	// Initialize Fiber instance
	// Similar to Echo's echo.New(), fiber.New() creates a new server instance
	app := fiber.New()

	// Define routes
	// Note: Fiber's Handler signature is func(*fiber.Ctx) error
	app.Get("/", func(c *fiber.Ctx) error {
		// Prepare response data
		response := models.ApiResponse{
			Code:    200,
			Message: "Query successful",
			Data: models.UserData{
				ID:    1,
				Email: "test@example.com",
				Role:  "Admin, fiber", // Marked as fiber version
			},
		}

		// Use c.Status().JSON() for output
		// Similar to Echo's c.JSON(), Fiber provides a convenient JSON response method
		return c.Status(http.StatusOK).JSON(response)
	})

	// Create user route
	// This route demonstrates how to parse request body into CreateUserRequest
	app.Post("/users", func(c *fiber.Ctx) error {
		// Declare request model
		var req models.CreateUserRequest

		// Parse request body JSON into struct
		if err := c.BodyParser(&req); err != nil {
			return c.Status(http.StatusBadRequest).JSON(models.ApiResponse{
				Code:    400,
				Message: "Invalid request body",
				Data:    nil,
			})
		}

		// Simulate created user data
		// In a real project, password is usually hashed and not returned directly
		createdUser := models.UserData{
			ID:    2,
			Email: req.Email,
			Role:  req.Role,
		}

		// Return standard response
		return c.Status(http.StatusCreated).JSON(models.ApiResponse{
			Code:    201,
			Message: "User created successfully",
			Data:    createdUser,
		})
	})

	// Startup message
	fmt.Println("Server started successfully! Please visit: http://localhost:8080")

	// Start server
	// app.Listen is equivalent to Echo's e.Start
	if err := app.Listen(":8080"); err != nil {
		panic(err)
	}
}
