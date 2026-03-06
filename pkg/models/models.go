package models

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
