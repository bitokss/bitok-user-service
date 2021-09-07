package controllers

import (
	"github.com/bitokss/bitok-user-service/domains/v1"
	"github.com/bitokss/bitok-user-service/services/v1"
	"github.com/bitokss/bitok-user-service/utils"
	"github.com/labstack/echo/v4"
)

var (
	ProfileController profileControllerInterface = &profileController{}
)

type profileControllerInterface interface {
	Update(c echo.Context) error
	Find(c echo.Context) error
}

type profileController struct{}

func (p profileController) Update(c echo.Context) error {
	uid, err := utils.ValidateAndCastToInt(c.Param("user_id"))
	if err != nil {
		return c.JSON(err.Status(), err)
	}
	user := new(domains.ProfileRequest)
	err = utils.ValidateAndBind(c, user)
	if err != nil {
		return c.JSON(err.Status(), err)
	}
	resp, err := services.ProfilesService.Update(uint(uid), *user)
	if err != nil {
		return c.JSON(err.Status(), err)
	}
	return c.JSON(resp.Status(), resp)
}

func (p profileController) Find(c echo.Context) error {
	uid, err := utils.ValidateAndCastToInt(c.Param("user_id"))
	if err != nil {
		return c.JSON(err.Status(), err)
	}
	resp, err := services.ProfilesService.Find(uint(uid))
	if err != nil {
		return c.JSON(err.Status(), err)
	}
	return c.JSON(resp.Status(), resp)
}
