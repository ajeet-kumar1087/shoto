package user

import (
	"gorm.io/gorm"
)

// User represents a user of the application.
type User struct {
	gorm.Model        // Add fields `ID`, `CreatedAt`, `UpdatedAt`, `DeletedAt`
	Email      string `gorm:"unique;not null"`
	Password   string `gorm:"not null"`
	FirstName  string
	LastName   string
}

type UserResponse struct {
	ID        string `json:"id"`
	FirstName string `json:"name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

type CreateUserRequest struct {
	FirstName string `json:"name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=8"`
}
