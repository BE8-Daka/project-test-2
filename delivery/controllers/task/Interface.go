package task

import "github.com/labstack/echo/v4"

type TaskController interface {
	Insert() echo.HandlerFunc
	GetAll() echo.HandlerFunc
	Update() echo.HandlerFunc
	Delete() echo.HandlerFunc
	UpdateStatus() echo.HandlerFunc
	GetTaskbyID() echo.HandlerFunc
}