package user

import "gorm.io/gorm"

func SaveUser(db *gorm.DB, user *User) error {
	// Create a new user
	if err := db.Create(&user).Error; err != nil {
		return err
	}
	return nil
}
