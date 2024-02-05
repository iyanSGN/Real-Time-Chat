package controller

import "github.com/labstack/echo/v4"

type Controller interface {
	GetAll(c echo.Context) error
	GetUserByID(c echo.Context) error
}