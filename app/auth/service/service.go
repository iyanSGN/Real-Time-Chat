package service

import (
	"realtime/app/auth"

	"github.com/labstack/echo/v4"
)

type Service interface {
	GetAll(c echo.Context) ([]auth.UserResponseDTO, error)
}
