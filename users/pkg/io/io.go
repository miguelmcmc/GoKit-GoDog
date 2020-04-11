package io

// User model for use service
type User struct {
	ID    string `json:"id,omitempty"`
	Name  string `json:"name" validate:"required"`
	Email string `json:"email"`
}
