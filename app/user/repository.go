package user

import "gorm.io/gorm"

type userRepository struct {
	db *gorm.DB
}

// Constructor for userRepository.
func NewRepository(db *gorm.DB) Repository {
	return &userRepository{db}
}

// GetUser implements Repository.
func (ur *userRepository) GetUser(id string) (*User, error) {
	var user User
	err := ur.db.First(&user, id)
	if err.Error != nil {
		return nil, err.Error
	}
	return &user, nil

}

// SaveUser implements Repository.
func (ur *userRepository) SaveUser(user *User) (*User, error) {
	if err := ur.db.Create(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// Update implements Repository.
func (ur *userRepository) UpdateUser(user *User) (*User, error) {
	if err := ur.db.Save(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// Delete implements Repository.
func (ur *userRepository) DeleteUser(id string) error {
	panic("unimplemented")
}
