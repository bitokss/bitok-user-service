package controllers

import "github.com/labstack/echo/v4"

var (
	RolesController rolesControllerInterface = &rolesController{}
)

type rolesControllerInterface interface {
	Create(c echo.Context) error
	FindAll(c echo.Context) error
	Find(c echo.Context) error
	Update(c echo.Context) error
	Delete(c echo.Context) error
}

type rolesController struct{}

func (r rolesController) Create(c echo.Context) error {
	panic("implement me")
}

func (r rolesController) FindAll(c echo.Context) error {
	panic("implement me")
}

func (r rolesController) Find(c echo.Context) error {
	panic("implement me")
}

func (r rolesController) Update(c echo.Context) error {
	panic("implement me")
}

func (r rolesController) Delete(c echo.Context) error {
	panic("implement me")
}
