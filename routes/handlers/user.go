package routes

import (
	"realtime/app/auth/controller"
	"realtime/app/auth/repository"
	"realtime/app/auth/service"
	"realtime/pkg/database"

	"github.com/labstack/echo/v4"
)

type handlerUsers struct {
	Controller controller.Controller
}


func UserHandler() *handlerUsers {
	s := service.NewService(database.DBManager(), repository.NewRepository())

	return &handlerUsers{
		Controller: controller.NewController(s),
	}
}

func (h *handlerUsers) Route(g *echo.Group) {
	g.GET("", h.Controller.GetAll)
	g.GET("/:id", h.Controller.GetUserByID)
	g.POST("", controller.CreateUser)
	g.PUT("/:id", controller.UpdateUser)
	g.DELETE("/:id", controller.DeleteUser)
}