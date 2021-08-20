package controllers

import "github.com/labstack/echo/v4"

var (
	PermissionsController permissionsControllerInterface = &permissionsController{}
)

type permissionsControllerInterface interface {
	Create(c echo.Context) error
	FindAll(c echo.Context) error
	Find(c echo.Context) error
	Update(c echo.Context) error
	Delete(c echo.Context) error
}

type permissionsController struct{}

func (p permissionsController) Create(c echo.Context) error {
	panic("implement me")
}

func (p permissionsController) FindAll(c echo.Context) error {
	panic("implement me")
}

func (p permissionsController) Find(c echo.Context) error {
	panic("implement me")
}

func (p permissionsController) Update(c echo.Context) error {
	panic("implement me")
}

func (p permissionsController) Delete(c echo.Context) error {
	panic("implement me")
}
