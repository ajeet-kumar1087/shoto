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
