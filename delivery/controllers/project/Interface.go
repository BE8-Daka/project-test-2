package project

import "github.com/labstack/echo/v4"

type ProjectController interface {
	Insert() echo.HandlerFunc
	GetAll() echo.HandlerFunc
}