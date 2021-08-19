package app

import (
	"github.com/bitokss/bitok-user-service/domains/v1"
	repositories "github.com/bitokss/bitok-user-service/repositories/postgres/v1"
	"github.com/labstack/echo/v4"
)

var (
	e *echo.Echo
)

func StartApp(port string) {
	e = echo.New()
	urlMapper()
	db := repositories.PostgresInit()
	err := db.AutoMigrate(&domains.User{},&domains.Level{})
	if err != nil {
		e.Logger.Error(err)
	}
	e.Logger.Error(e.Start(port))
}