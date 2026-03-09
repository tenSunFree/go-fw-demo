package main

import (
	// Format I/O (for printing text)
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-mizu/go-fw/pkg/models"
)

func main() {
	// r := gin.New()
	// It is recommended to use gin.Default() instead of gin.New()
	// Default() will automatically add a Logger (request log) and a Recovery mechanism (to prevent program crashes).
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		// Prepare the data to be sent back
		response := models.ApiResponse{
			Code:    200, // customize the business logic status code here
			Message: "Query successful",
			Data: models.UserData{
				ID:    1,
				Email: "test@example.com",
				Role:  "Admin, gin",
			},
		}
		// Use c.JSON() to output
		// The first parameter is the HTTP Status Code (200 OK)
		// The second parameter is the structure you want to convert to JSON (Struct)
		c.JSON(http.StatusOK, response)
		// c.String(http.StatusOK, "hello, world!")
	})
	fmt.Println("Server started successfully! Please visit: http://localhost:8080")
	r.Run(":8080")
}
