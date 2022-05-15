package routes

import (
	"project-test/delivery/controllers/project"
	"project-test/delivery/controllers/task"
	"project-test/delivery/controllers/user"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func UserPath(e *echo.Echo, c user.UserController) {
	e.POST("/users/register", c.Insert())
	e.POST("/users/login", c.Login())
	
	auth := e.Group("/users", middleware.JWT([]byte("$4dm!n$")))
	auth.GET("/profile", c.GetbyID())
	auth.PUT("/profile", c.Update())
	auth.DELETE("/profile", c.Delete())
}

func TaskPath(e *echo.Echo, c task.TaskController) {
	auth := e.Group("/tasks", middleware.JWT([]byte("$4dm!n$")))
	auth.POST("", c.Insert())
}

func ProjectPath(e *echo.Echo, c project.ProjectController) {
	auth := e.Group("/projects", middleware.JWT([]byte("$4dm!n$")))
	auth.POST("", c.Insert())
	auth.GET("", c.GetAll())
}