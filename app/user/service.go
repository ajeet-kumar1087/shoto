package user

import (
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo}
}

// CreateUser implements Service.
func (s *service) CreateUser(user *User) (*User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("failed to hash password")
	}

	user.Password = string(hashedPassword)

	return s.repo.SaveUser(user)
}

// GetUser implements Service.
func (s *service) GetUser(id string) (*User, error) {
	return s.repo.GetUser(id)
}

// UpdateUser implements Service.
func (s *service) UpdateUser(id string, req *User) (*User, error) {
	// Find existing user
	existingUser, err := s.repo.GetUser(id)
	if err != nil {
		return nil, err
	}

	// Update only provided fields
	if req.FirstName != "" {
		existingUser.FirstName = req.FirstName
	}
	if req.LastName != "" {
		existingUser.LastName = req.LastName
	}
	if req.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			return nil, err
		}
		existingUser.Password = string(hashedPassword)
	}

	// Save updates
	updatedUser, err := s.repo.UpdateUser(existingUser)
	if err != nil {
		return nil, fmt.Errorf("failed to update user: %w", err)
	}

	return updatedUser, nil
}

// DeleteUser implements Service.
func (s *service) DeleteUser(id string) error {
	panic("unimplemented")
}
