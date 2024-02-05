package service

import (
	"realtime/app/auth"

	"github.com/labstack/echo/v4"
)

type Service interface {
	GetAll(c echo.Context) ([]auth.UserResponseDTO, error)
	GetUserByID(c echo.Context, id uint) (auth.UserResponseDTO, error)
}
