package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"

	"project-test/config"
	"project-test/delivery/controllers/project"
	"project-test/delivery/controllers/task"
	"project-test/delivery/controllers/user"
	"project-test/delivery/middlewares"
	"project-test/delivery/routes"
	repoProject "project-test/repository/project"
	repoTask "project-test/repository/task"
	repoUser "project-test/repository/user"
)
func main() {
	setting := config.InitConfig()
	db := config.InitDB(*setting)
	config.AutoMigrate(db)

	e := echo.New()
	
	middlewares.General(e)
	routes.UserPath(e, user.NewUserController(repoUser.NewUserModel(db), validator.New()))
	routes.TaskPath(e, task.NewTaskController(repoTask.NewTaskModel(db), validator.New()))
	routes.ProjectPath(e, project.NewProjectController(repoProject.NewProjectModel(db), validator.New()))

	e.Start(":8000")
}