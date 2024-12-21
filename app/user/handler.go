package user

import (
	"net/http"

	"github.com/labstack/echo"
)

func CreateUser(c echo.Context) error {
	// Create a new user
	return c.JSON(http.StatusOK, "User created")
}
