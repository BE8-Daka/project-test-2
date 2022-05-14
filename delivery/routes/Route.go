package routes

import (
	user "project-test/delivery/controllers/user"

	"github.com/labstack/echo/v4"
)

func UserPath(e *echo.Echo, c user.UserController) {
	e.POST("/users/register", c.Insert())
	e.POST("/users/login", c.Login())
}