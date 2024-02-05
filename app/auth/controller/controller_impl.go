package controller

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"net/http"
	"path/filepath"
	"realtime/app/auth"
	"realtime/app/auth/repository"
	"realtime/app/auth/service"
	"realtime/models"
	helpers "realtime/pkg/helper"
	"realtime/pkg/response"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/nfnt/resize"
)

type controllerImpl struct {
	Service service.Service
}

func NewController(Service service.Service) Controller {
	return &controllerImpl{
		Service: Service,
	}
}

func (co *controllerImpl) GetAll(c echo.Context) error {
	result, err := co.Service.GetAll(c)
	if err != nil {
		return err
	}

	return response.SuccessResponse(c, http.StatusOK, "Success Get All User Chatin'g", result)
}

func (co *controllerImpl) GetUserByID(c echo.Context) error {
	userID := c.Param("id")
	id, err := strconv.ParseUint(userID, 10, 64)
	if err != nil {
		return response.ErrorResponse(c, response.BuildError(response.ErrNotFound, err))
	}

	result, err := co.Service.GetUserByID(c, uint(id))
	if err != nil {
		return response.ErrorResponse(c, err)
	}

	return response.SuccessResponse(c, http.StatusOK, "Success Get user by id", result)
}

func CreateUser(c echo.Context) error {
	data := auth.UserRequest{}

	if err := c.Bind(&data); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := helpers.HashPassword(&data); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	var base64Image string
	var MainPhoto []byte

	file, err := c.FormFile("main_photo")
	if err != nil {
		fmt.Println("no image uploaded")

	} else {
		src, err := file.Open()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message":     "Failed to open uploaded image",
				"error":       err.Error(),
				"status_code": http.StatusInternalServerError,
			})
		}
		defer src.Close()

		MainPhoto, err = io.ReadAll(src)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message":     "Failed to read uploaded image",
				"error":       err.Error(),
				"status_code": http.StatusInternalServerError,
			})
		}

		newWidth := 600
		img, _, err := image.Decode(bytes.NewReader(MainPhoto))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message":     "Failed to decode uploaded image",
				"error":       err.Error(),
				"status_code": http.StatusInternalServerError,
			})
		}
		resizedImg := resize.Resize(uint(newWidth), 0, img, resize.Lanczos3)
		var resizedImgBuffer bytes.Buffer
		switch filepath.Ext(file.Filename) {
		case ".jpg", ".jpeg":
			if err := jpeg.Encode(&resizedImgBuffer, resizedImg, nil); err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]interface{}{
					"message":     "Failed to encode resized image",
					"error":       err.Error(),
					"status_code": http.StatusInternalServerError,
				})
			}
		case ".png":
			if err := png.Encode(&resizedImgBuffer, resizedImg); err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]interface{}{
					"message":     "Failed to encode resized png image",
					"error":       err.Error(),
					"status_code": http.StatusInternalServerError,
				})
			}

		default:
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message":     "Unknown image format",
				"status_code": http.StatusBadRequest,
			})
		}

		base64Image = base64.StdEncoding.EncodeToString(resizedImgBuffer.Bytes())
	}

	createdUser, err := repository.CreateUser(data, base64Image)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message":     "Failed to create user",
			"error":       err.Error(),
			"status_code": http.StatusInternalServerError,
		})
	}

	fmt.Println(createdUser)

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message":     "User created successfully",
		"status_code": http.StatusCreated,
	})
}

func UpdateUser(c echo.Context) error {
	var user models.User
	userID := c.Param("id")
	id, err := strconv.ParseUint(userID, 10, 64)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, fmt.Errorf("invalid user id"))
	}

	userUpdate := auth.UserRequest{}
	if err := c.Bind(&userUpdate); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if userUpdate.Password != "" {
		if err := helpers.HashPassword(&userUpdate); err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
	}

	var base64Image string
	var MainPhoto []byte

	file, err := c.FormFile("main_photo")
	if err != nil || file == nil {
		
		base64Image = user.MainPhoto

		err = repository.UpdateUser(userUpdate, base64Image, uint(id))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

	} else {
		src, err := file.Open()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message":     "Failed to open uploaded image",
				"error":       err.Error(),
				"status_code": http.StatusInternalServerError,
			})
		}
		defer src.Close()

		MainPhoto, err = io.ReadAll(src)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message":     "Failed to read uploaded image",
				"error":       err.Error(),
				"status_code": http.StatusInternalServerError,
			})
		}

		newWidth := 600
		img, _, err := image.Decode(bytes.NewReader(MainPhoto))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message":     "Failed to decode uploaded image",
				"error":       err.Error(),
				"status_code": http.StatusInternalServerError,
			})
		}
		resizedImg := resize.Resize(uint(newWidth), 0, img, resize.Lanczos3)
		var resizedImgBuffer bytes.Buffer
		switch filepath.Ext(file.Filename) {
		case ".jpg", ".jpeg":
			if err := jpeg.Encode(&resizedImgBuffer, resizedImg, nil); err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]interface{}{
					"message":     "Failed to encode resized image",
					"error":       err.Error(),
					"status_code": http.StatusInternalServerError,
				})
			}
		case ".png":
			if err := png.Encode(&resizedImgBuffer, resizedImg); err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]interface{}{
					"message":     "Failed to encode resized png image",
					"error":       err.Error(),
					"status_code": http.StatusInternalServerError,
				})
			}

		default:
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message":     "Unknown image format",
				"status_code": http.StatusBadRequest,
			})
		}

		base64Image = base64.StdEncoding.EncodeToString(resizedImgBuffer.Bytes())

		fmt.Println(userUpdate)

		err = repository.UpdateUser(userUpdate, base64Image, uint(id))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":     "User Updated Successfully.",
		"status_code": http.StatusOK,
	})

}

func DeleteUser(c echo.Context) error {
	userID := c.Param("id")
	id, err :=strconv.ParseUint(userID, 10,64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err = repository.DeleteUser(uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "selected user has been deleted",
		"status_code" : http.StatusOK,
	})

}
