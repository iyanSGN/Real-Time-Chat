package main

import (
	"realtime/pkg/database"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	database.EnvInit()
	database.Init("postgresql")
	database.Migrate()

	e.Start(":1234")
}
