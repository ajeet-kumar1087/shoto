package user

import "github.com/labstack/echo"

type Repository interface {
	GetUser(id string) (*User, error)
	SaveUser(user *User) (*User, error)
	UpdateUser(user *User) (*User, error)
	DeleteUser(id string) error
}

type Service interface {
	GetUser(id string) (*User, error)
	CreateUser(user *User) (*User, error)
	UpdateUser(id string, user *User) (*User, error)
	DeleteUser(id string) error
}

type Handler interface {
	GetUser(c echo.Context) error
	CreateUser(c echo.Context) error
	UpdateUser(c echo.Context) error
	DeleteUser(c echo.Context) error
}
