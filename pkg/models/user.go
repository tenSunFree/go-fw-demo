package models

// Specific business data structure
type UserData struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
	Role  string `json:"role"`
}

// Request structure for creating a user
type CreateUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

// Request structure for updating a user
type UpdateUserRequest struct {
	Email string `json:"email"`
	Role  string `json:"role"`
}
