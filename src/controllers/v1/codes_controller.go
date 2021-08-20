package controllers

import "github.com/labstack/echo/v4"

var (
	CodesController codesControllerInterface = &codesController{}
)

type codesControllerInterface interface {
	Send(c echo.Context) error
	Verify(c echo.Context) error
}

type codesController struct{}

func (c2 codesController) Send(c echo.Context) error {
	panic("implement me")
}

func (c2 codesController) Verify(c echo.Context) error {
	panic("implement me")
}
