package user

import (
	"net/http"

	"github.com/labstack/echo"
)

type handler struct {
	service Service
}

func NewHandler(s Service) Handler {
	return &handler{s}
}

// CreateUser implements Handler.
func (h *handler) CreateUser(c echo.Context) error {
	var req CreateUserRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	u, err := h.service.CreateUser(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, u)
}

// GetUser implements Handler.
func (h *handler) GetUser(c echo.Context) error {
	id := c.Param("id")
	user, err := h.service.GetUser(id)
	if err != nil {
		return c.JSON(500, map[string]string{"error": err.Error()})
	}
	return c.JSON(200, user)
}

// DeleteUser implements Handler.
func (h *handler) DeleteUser(c echo.Context) error {
	panic("unimplemented")
}

// UpdateUser implements Handler.
func (h *handler) UpdateUser(c echo.Context) error {
	panic("unimplemented")
}
