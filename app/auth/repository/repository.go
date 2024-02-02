package repository

import (
	"realtime/models"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Repository interface {
	GetAll(c echo.Context, DB gorm.DB) ([]models.User, error)
}