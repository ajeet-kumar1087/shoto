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
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type CreateUserRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}
