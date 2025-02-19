package user

import "github.com/labstack/echo"

type UserHandler struct {
	Service UserService
}

func NewUserHandler(service UserService) *UserHandler {
	return &UserHandler{Service: service}
}

func (h *UserHandler) GetUser(c echo.Context) error {
	id := c.Param("id")
	user, err := h.Service.GetUserByID(id)
	if err != nil {
		return c.JSON(500, map[string]string{"error": err.Error()})
	}
	return c.JSON(200, user)
}

func (h *UserHandler) CreateUser(c echo.Context) error {

	return c.String(200, "Create user")
}
func (h *UserHandler) UpdateUser(c echo.Context) error {
	return c.String(200, "Update user")
}

func (h *UserHandler) DeleteUser(c echo.Context) error {
	return c.String(200, "Delete user")
}

func (h *UserHandler) HealthCheck(c echo.Context) error {
	return c.JSON(200, map[string]string{"status": "ok", "message": "User Health is okay okay ok"})
}
