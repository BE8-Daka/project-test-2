package user

import "github.com/labstack/echo/v4"

type UserController interface {
	Insert() echo.HandlerFunc
}