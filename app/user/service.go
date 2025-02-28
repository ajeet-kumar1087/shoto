package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo}
}

// CreateUser implements Service.
func (s *service) CreateUser(req CreateUserRequest) (*User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("failed to hash password")
	}
	user := &User{
		Email:    req.Email,
		Password: string(hashedPassword),
	}
	return user, s.repo.SaveUser(user)
}

// DeleteUser implements Service.
func (s *service) DeleteUser(id string) error {
	panic("unimplemented")
}

// GetUser implements Service.
func (s *service) GetUser(id string) (*User, error) {
	panic("unimplemented")
}

// UpdateUser implements Service.
func (s *service) UpdateUser(id string, user *User) error {
	panic("unimplemented")
}
