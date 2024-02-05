package main

import (
	"realtime/pkg/database"
	"realtime/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	database.EnvInit()
	database.Init("postgresql")
	database.Migrate()

	routes.RouteInit(e.Group("api"))

	e.Logger.Fatal(e.Start(":" + database.Get("APP_PORT")))
}
