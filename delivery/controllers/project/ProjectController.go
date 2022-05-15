package project

import (
	"net/http"
	"project-test/entity"

	"github.com/labstack/echo/v4"

	"project-test/delivery/middlewares"
	"project-test/delivery/views/request"
	"project-test/delivery/views/response"

	"github.com/go-playground/validator/v10"

	repo "project-test/repository/project"
)

type projectController struct {
	Connect 	repo.ProjectModel
	Validate 	*validator.Validate
}

func NewProjectController(db repo.ProjectModel, valid *validator.Validate) *projectController {
	return &projectController{
		Connect: db,
		Validate: valid,
	}
}

func (c *projectController) Insert() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		user_id := middlewares.ExtractTokenUserId(ctx)
		var request request.InsertProject

		if err := ctx.Bind(&request); err != nil {
			return ctx.JSON(http.StatusBadRequest, response.StatusBadRequestBind(err))
		}

		if err := c.Validate.Struct(request); err != nil {
			return ctx.JSON(http.StatusBadRequest, response.StatusBadRequestRequired(err))
		}

		project := entity.Project{
			Name: request.Name,
			UserID: uint(user_id),
		}

		result, err := c.Connect.Insert(&project)
		if err != nil {
			return ctx.JSON(http.StatusBadRequest, response.StatusBadRequestDuplicate(err))
		}

		return ctx.JSON(http.StatusCreated, response.StatusCreated(result))
	}
}