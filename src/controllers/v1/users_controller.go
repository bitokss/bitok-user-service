package controllers

import (
	"github.com/labstack/echo/v4"
)

var UserController userControllerInterface = &userController{}

type userControllerInterface interface {
	Find(c echo.Context) error
}

type userController struct {}

func (u *userController) Find(c echo.Context) error {
	return c.JSON(200 , "user :)")
}

