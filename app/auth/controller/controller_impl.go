package controller

import (
	"net/http"
	"realtime/app/auth/service"
	"realtime/pkg/response"

	"github.com/labstack/echo/v4"
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