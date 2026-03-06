package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
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

func main() {
	// r := gin.New()
	// It is recommended to use gin.Default() instead of gin.New()
	// Default() will automatically add a Logger (request log) and a Recovery mechanism (to prevent program crashes).
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		// Prepare the data to be sent back
		response := ApiResponse{
			Code:    200, // customize the business logic status code here
			Message: "Query successful",
			Data: UserData{
				ID:    1,
				Email: "test@example.com",
				Role:  "Admin",
			},
		}
		// Use c.JSON() to output
		// The first parameter is the HTTP Status Code (200 OK)
		// The second parameter is the structure you want to convert to JSON (Struct)
		c.JSON(http.StatusOK, response)
		// c.String(http.StatusOK, "hello, world!")
	})
	r.Run(":8080")
}
