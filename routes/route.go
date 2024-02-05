package routes

import (
	"net/http"
	"realtime/pkg/database"
	"realtime/routes/handlers"

	"github.com/labstack/echo/v4"
)

func RouteInit(g *echo.Group) {
	g.GET("", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome to "+ database.Get("APP_NAME")+"! version "+database.Get("APP_VERSION")+" in mode "+database.Get("ENV"))
	})

	routes.UserHandler().Route(g.Group("/user"))

}