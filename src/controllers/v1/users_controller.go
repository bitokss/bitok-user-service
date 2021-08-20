package controllers

import (
	"github.com/labstack/echo/v4"
)

var (
	UserController userControllerInterface = &userController{}
)

type userControllerInterface interface {
	Create(c echo.Context) error
	FindAll(c echo.Context) error
	Find(c echo.Context) error
	Update(c echo.Context) error
	Delete(c echo.Context) error
	Register(c echo.Context) error
	Login(c echo.Context) error
	FindByToken(c echo.Context) error
	FindByUsername(c echo.Context) error
	ResetPassword(c echo.Context) error
	TickRequest(c echo.Context) error
}

type userController struct{}

func (u *userController) Create(c echo.Context) error {
	panic("implement me")
}

func (u *userController) FindAll(c echo.Context) error {
	panic("implement me")
}

func (u *userController) Update(c echo.Context) error {
	panic("implement me")
}

func (u *userController) Delete(c echo.Context) error {
	panic("implement me")
}

func (u *userController) Register(c echo.Context) error {
	panic("implement me")
}

func (u *userController) Login(c echo.Context) error {
	panic("implement me")
}

func (u *userController) FindByToken(c echo.Context) error {
	panic("implement me")
}

func (u *userController) FindByUsername(c echo.Context) error {
	panic("implement me")
}

func (u *userController) ResetPassword(c echo.Context) error {
	panic("implement me")
}

func (u *userController) TickRequest(c echo.Context) error {
	panic("implement me")
}

func (u *userController) Find(c echo.Context) error {
	panic("implement me")
}
