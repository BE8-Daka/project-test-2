package user

import "github.com/labstack/echo/v4"

type UserController interface {
	Insert() echo.HandlerFunc
	Login() echo.HandlerFunc
	GetbyID() echo.HandlerFunc
	Update() echo.HandlerFunc
}