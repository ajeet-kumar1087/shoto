package user

import (
	"fmt"
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
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request format",
		})
	}

	// Convert to User model
	user := &User{
		Email:     req.Email,
		Password:  req.Password,
		FirstName: req.FirstName,
		LastName:  req.LastName,
	}

	// Save using service layer
	savedUser, err := h.service.CreateUser(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	// Convert to response
	response := UserResponse{
		ID:        fmt.Sprint(savedUser.ID),
		Email:     savedUser.Email,
		FirstName: savedUser.FirstName,
		LastName:  savedUser.LastName,
	}

	return c.JSON(http.StatusCreated, response)
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

// UpdateUser implements Handler.
func (h *handler) UpdateUser(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "ID required"})
	}
	var req UpdateUserRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request format"})
	}

	user := &User{
		Password:  req.Password,
		FirstName: req.FirstName,
		LastName:  req.LastName,
	}
	// Call service layer
	updatedUser, err := h.service.UpdateUser(id, user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}
	response := UserResponse{
		ID:        fmt.Sprint(updatedUser.ID),
		Email:     updatedUser.Email,
		FirstName: updatedUser.FirstName,
		LastName:  updatedUser.LastName,
	}
	return c.JSON(http.StatusOK, response)

}

// DeleteUser implements Handler.
func (h *handler) DeleteUser(c echo.Context) error {
	panic("unimplemented")
}
