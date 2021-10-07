package controllers

import (
	"github.com/bitokss/bitok-user-service/src/domains/v1"
	"github.com/bitokss/bitok-user-service/src/services/v1"
	"github.com/bitokss/bitok-user-service/src/utils"
	"github.com/labstack/echo/v4"
)

var (
	CodesController codesControllerInterface = &codesController{}
)

type codesControllerInterface interface {
	Send(c echo.Context) error
	Verify(c echo.Context) error
}

type codesController struct{}

func (*codesController) Send(c echo.Context) error {
	body := new(domains.CodeRequest)
	if err := utils.ValidateAndBind(c, body); err != nil {
		return c.JSON(err.Status(), err)
	}
	if err := utils.IsValidCodeType(body.Type); err != nil {
		return c.JSON(err.Status(), err)
	}
	resp, err := services.CodesService.Send(*body)
	if err != nil {
		return c.JSON(err.Status(), err)
	}
	return c.JSON(resp.Status(), resp)
}

func (*codesController) Verify(c echo.Context) error {
	body := new(domains.VerifyRequest)
	if err := utils.ValidateAndBind(c, body); err != nil {
		return c.JSON(err.Status(), err)
	}
	if err := utils.IsValidCodeType(body.Type); err != nil {
		return c.JSON(err.Status(), err)
	}
	resp, err := services.CodesService.Verify(*body)
	if err != nil {
		return c.JSON(err.Status(), err)
	}
	return c.JSON(resp.Status(), resp)
}
