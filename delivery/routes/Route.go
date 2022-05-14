package routes

import (
	user "project-test/delivery/controllers/user"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func UserPath(e *echo.Echo, c user.UserController) {
	e.POST("/users/register", c.Insert())
	e.POST("/users/login", c.Login())
	
	auth := e.Group("/users", middleware.JWT([]byte("$4dm!n$")))
	auth.GET("/profile", c.GetbyID())
	auth.PUT("/profile", c.Update())
}