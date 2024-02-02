package repository

import (
	"realtime/models"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type repositoryImpl struct {
}

func NewRepository() Repository {
	return &repositoryImpl{}
}

func (r *repositoryImpl) GetAll(c echo.Context, DB gorm.DB) ([]models.User, error) {
	var user []models.User

	if err := DB.Find(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}