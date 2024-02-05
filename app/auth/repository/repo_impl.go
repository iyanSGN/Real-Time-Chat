package repository

import (
	"fmt"
	"realtime/app/auth"
	"realtime/models"
	"realtime/pkg/database"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type repositoryImpl struct {
}

func NewRepository() Repository {
	return &repositoryImpl{}
}

func (r *repositoryImpl) GetAll(c echo.Context, DB *gorm.DB) ([]models.User, error) {
    var users []models.User

    if err := DB.Find(&users).Error; err != nil {
        return nil, err
    }

    return users, nil
}

func (r *repositoryImpl) GetUserByID(c echo.Context, DB gorm.DB, id uint) (models.User, error) {
	var user models.User

	if err := DB.Where("id = ?", id).Find(&user).Error; 
	
	err != nil {
		return user, err
	}

	return user, nil
}

func CreateUser(req auth.UserRequest, mainPhoto string) (auth.UserRequest, error) {
	db:= database.DBManager()

	user := models.User{
		UserName: req.Name,
		Password: req.Password,
		Phone: req.Phone,
		MainPhoto: mainPhoto,
	}

	result := db.Create(&user)
	if result.Error != nil {
		return req, fmt.Errorf("error creating user: %v", result.Error)
	}

	return req, nil
}

func UpdateUser(req auth.UserRequest, mainPhoto string, id uint) ( error) {
	db := database.DBManager()

	var user models.User
	result := db.First(&user,id)
	if result.Error != nil {
		return  fmt.Errorf("error creating user: %v", result.Error)
	}

	if req.Name != "" {
		user.UserName = req.Name
	}

	if req.Password != "" {
		user.Password = req.Password
	}

	if req.Email != "" {
		user.Email = req.Email
	}

	if req.Phone != "" {
		user.Phone = req.Phone
	}

	updatedUser := db.Save(&user)
	if updatedUser != nil {
		return fmt.Errorf("error updating user: %w", updatedUser.Error)
	}

	return nil
}

func DeleteUser(id uint) error {
	db := database.DBManager()

	var user models.User
	result := db.First(&user, id)
	if result.Error != nil {
		return fmt. Errorf("error deleting user: %w", result.Error)
	}

	deleteUser := db.Delete(&user)
	if deleteUser.Error != nil {
		return fmt.Errorf("error deleting user: %w", deleteUser.Error)
	}
	return nil
}