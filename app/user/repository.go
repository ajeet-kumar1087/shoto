package user

import "gorm.io/gorm"

type UserRepository interface {
	FindByID(id string) (*User, error)
	SaveUser(user *User) (*User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (ur *userRepository) FindByID(id string) (*User, error) {
	var user User
	if err := ur.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (ur *userRepository) SaveUser(user *User) (*User, error) {
	if err := ur.db.Create(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
