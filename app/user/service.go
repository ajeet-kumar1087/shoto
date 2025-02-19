package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	Repo UserRepository
}

type UserService interface {
	GetUserByID(id string) (*User, error)
	CreateUser(req CreateUserRequest) (*User, error)
}

func NewUserService(repo UserRepository) UserService {
	return &userService{Repo: repo}
}

func (s *userService) GetUserByID(id string) (*User, error) {
	return s.Repo.FindByID(id)
}

func (s *userService) CreateUser(req CreateUserRequest) (*User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("failed to hash password")
	}
	user := &User{
		Email:    req.Email,
		Password: string(hashedPassword),
	}
	return s.Repo.SaveUser(user)
}
