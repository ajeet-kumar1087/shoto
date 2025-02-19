package user

import "github.com/labstack/echo"

func Routes(e *echo.Echo, handler *UserHandler) {
	e.GET("/health", handler.HealthCheck)
	e.GET("/user/:id", handler.GetUser)
	e.POST("/user", handler.CreateUser)
	e.PUT("/user/:id", handler.UpdateUser)
	e.DELETE("/user/:id", handler.DeleteUser)
}
