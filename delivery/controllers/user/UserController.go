package user

import (
	"net/http"
	"project-test/entity"

	"github.com/labstack/echo/v4"

	"project-test/delivery/views/request"
	"project-test/delivery/views/response"

	"github.com/go-playground/validator/v10"

	repo "project-test/repository/user"
)

type userController struct {
	Connect 	repo.UserModel
	Validate 	*validator.Validate
}

func NewUserController(db repo.UserModel, valid *validator.Validate) *userController {
	return &userController{
		Connect: db,
		Validate: valid,
	}
}

func (c *userController) Insert() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var request request.InsertUser

		if err := ctx.Bind(&request); err != nil {
			return ctx.JSON(http.StatusBadRequest, response.StatusBadRequestBind(err))
		}

		if err := c.Validate.Struct(request); err != nil {
			return ctx.JSON(http.StatusBadRequest, response.StatusBadRequestRequired(err))
		}

		user := entity.User{
			Name: request.Name,
			Username: request.Username,
			NoHp: request.NoHp,
			Email: request.Email,
			Password: request.Password,
		}


		result, err := c.Connect.Insert(&user)
		if err != nil {
			return ctx.JSON(http.StatusBadRequest, response.StatusBadRequestDuplicate(err))
		}

		return ctx.JSON(http.StatusCreated, response.StatusCreated(result))
	}
}