package models

// Standard API response structure
type ApiResponse struct {
	Code    int         `json:"code"`    // Custom status code
	Message string      `json:"message"` // Message for frontend display
	Data    interface{} `json:"data"`    // Actual data content, can be anything
}

// Standard error response structure
type ErrorResponse struct {
	Code    int    `json:"code"`    // Error code
	Message string `json:"message"` // Error message
}

// Standard pagination information
type Pagination struct {
	Page      int `json:"page"`       // Current page
	PageSize  int `json:"page_size"`  // Number of items per page
	Total     int `json:"total"`      // Total number of items
	TotalPage int `json:"total_page"` // Total number of pages
}

// Response structure for paginated data
type PageResponse struct {
	Code       int         `json:"code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
	Pagination Pagination  `json:"pagination"`
}
