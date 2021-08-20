package controllers

import "github.com/labstack/echo/v4"

var (
	ProfileController profileControllerInterface = &profileController{}
)

type profileControllerInterface interface {
	CreateOrUpdate(c echo.Context) error
	Find(c echo.Context) error
}

type profileController struct{}

func (p profileController) CreateOrUpdate(c echo.Context) error {
	panic("implement me")
}

func (p profileController) Find(c echo.Context) error {
	panic("implement me")
}
