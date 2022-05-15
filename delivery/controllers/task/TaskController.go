package task

import (
	"net/http"
	"project-test/entity"
	"strconv"

	"github.com/labstack/echo/v4"

	"project-test/delivery/middlewares"
	"project-test/delivery/views/request"
	"project-test/delivery/views/response"

	"github.com/go-playground/validator/v10"

	repo "project-test/repository/task"
)

type taskController struct {
	Connect 	repo.TaskModel
	Validate 	*validator.Validate
}

func NewTaskController(db repo.TaskModel, valid *validator.Validate) *taskController {
	return &taskController{
		Connect: db,
		Validate: valid,
	}
}

func (c *taskController) Insert() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		user_id := middlewares.ExtractTokenUserId(ctx)
		var request request.InsertTask

		if err := ctx.Bind(&request); err != nil {
			return ctx.JSON(http.StatusBadRequest, response.StatusBadRequestBind(err))
		}

		if err := c.Validate.Struct(request); err != nil {
			return ctx.JSON(http.StatusBadRequest, response.StatusBadRequestRequired(err))
		}

		task := entity.Task{
			Name: request.Name,
			Status: true,
			UserID: uint(user_id),
			ProjectID: request.ProjectID,
		}

		result := c.Connect.Insert(&task)

		return ctx.JSON(http.StatusCreated, response.StatusCreated(result))
	}
}

func (c *taskController) GetAll() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		user_id := middlewares.ExtractTokenUserId(ctx)

		results := c.Connect.GetAll(uint(user_id))

		return ctx.JSON(http.StatusOK, response.StatusOK("get all data", results))
	}
}

func (c *taskController) Update() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		user_id := middlewares.ExtractTokenUserId(ctx)
		id, _ := strconv.Atoi(ctx.Param("id"))
		var request request.UpdateTask

		if !c.Connect.CheckExist(uint(id), uint(user_id)) {
			return ctx.JSON(http.StatusForbidden, response.StatusForbidden())
		}

		if err := ctx.Bind(&request); err != nil {
			return ctx.JSON(http.StatusBadRequest, response.StatusBadRequestBind(err))
		}

		task := entity.Task{
			Name: request.Name,
			ProjectID: request.ProjectID,
		}

		result, err := c.Connect.Update(uint(id), &task)
		if err != nil {
			return ctx.JSON(http.StatusBadRequest, response.StatusBadRequest(err))
		}
		
		return ctx.JSON(http.StatusOK, response.StatusOK("updated", result))
	}
}