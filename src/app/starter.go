package app

import (
	"github.com/bitokss/bitok-user-service/constants"
	"github.com/bitokss/bitok-user-service/domains/v1"
	"github.com/bitokss/bitok-user-service/repositories/postgres/v1"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"net/http"
)

var (
	e *echo.Echo
)

type Validator struct {
	validator *validator.Validate
}

func (v *Validator) Validate(i interface{}) error {
	if err := v.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, constants.InvalidInputErr)
	}
	return nil
}

func StartApp(port string) {
	e = echo.New()
	// validate inputs using go-playground package
	e.Validator = &Validator{validator: validator.New()}
	urlMapper()
	// initialize postgres and get db instance
	db := repositories.PostgresInit()
	// autoMigrate will automatically create tables using domains
	err := db.AutoMigrate(&domains.Permission{}, &domains.Role{}, &domains.Level{} , &domains.User{}, &domains.Code{})
	if err != nil {
		e.Logger.Error(err)
	}
	// start echo server
	e.Logger.Error(e.Start(port))
}
