package controllers

import "github.com/labstack/echo/v4"

var (
	LevelsController levelsControllerInterface = &levelsController{}
)

type levelsControllerInterface interface {
	Create(c echo.Context) error
	FindAll(c echo.Context) error
	Find(c echo.Context) error
	Update(c echo.Context) error
	Delete(c echo.Context) error
}

type levelsController struct{}

func (l levelsController) Create(c echo.Context) error {
	panic("implement me")
}

func (l levelsController) FindAll(c echo.Context) error {
	panic("implement me")
}

func (l levelsController) Find(c echo.Context) error {
	panic("implement me")
}

func (l levelsController) Update(c echo.Context) error {
	panic("implement me")
}

func (l levelsController) Delete(c echo.Context) error {
	panic("implement me")
}
