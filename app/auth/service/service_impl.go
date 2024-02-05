package service

import (
	"fmt"
	"realtime/app/auth"
	"realtime/app/auth/repository"
	// "realtime/pkg/response"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type serviceImpl struct {
	DB *gorm.DB
	Repository repository.Repository
}

func NewService(DB *gorm.DB, Repository repository.Repository) Service {
	return &serviceImpl{
		DB: DB,
		Repository: Repository,
	}
}

func (s *serviceImpl) GetAll(c echo.Context) ([]auth.UserResponseDTO, error) {
    var userRes []auth.UserResponseDTO
    
    result, err := s.Repository.GetAll(c, s.DB)
    if err != nil {
        return nil, fmt.Errorf("error getting all users: %v", err)
    }

    for _, user := range result {
        userRes = append(userRes, user.ToResponse())
    }

    return userRes, nil
}

func (s *serviceImpl) GetUserByID(c echo.Context, id uint) (auth.UserResponseDTO, error) {
	var userRes auth.UserResponseDTO

	result, err := s.Repository.GetUserByID(c, *s.DB, id)
	if err != nil {
		return userRes, err
	}

	user := result.ToResponse()

	return user, nil
}