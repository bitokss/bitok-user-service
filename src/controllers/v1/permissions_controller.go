package controllers

import (
	"github.com/alidevjimmy/go-rest-utils/rest_response"
	"github.com/bitokss/bitok-user-service/constants"
	"github.com/bitokss/bitok-user-service/domains/v1"
	"github.com/bitokss/bitok-user-service/services/v1"
	"github.com/bitokss/bitok-user-service/utils"
	"github.com/labstack/echo/v4"
	"strconv"
)

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
	permission := new(domains.CreatePermissionsRequest)
	// validate permission request data and bind it on CreatePermissionsRequest struct
	err := utils.ValidateAndBind(c, permission)
	if err != nil {
		return c.JSON(err.Status(), err)
	}
	resp, err := services.PermissionsService.Create(*permission)
	if err != nil {
		return c.JSON(err.Status(), err)
	}
	return c.JSON(resp.Status(), resp)
}

func (p permissionsController) FindAll(c echo.Context) error {
	limit := 50
	offset := 0
	limitParam := c.Param("limit")
	offsetParam := c.Param("offset")
	// check if param not sent, setting values of default
	if limitParam != "" {
		l, err := strconv.Atoi(limitParam)
		if err != nil {
			restErr := rest_response.NewBadRequestError(constants.InvalidInputErr, nil)
			return c.JSON(restErr.Status(), restErr)
		}
		limit = l
	}
	if offsetParam != "" {
		o, err := strconv.Atoi(offsetParam)
		if err != nil {
			restErr := rest_response.NewBadRequestError(constants.InvalidInputErr, nil)
			return c.JSON(restErr.Status(), restErr)
		}
		offset = o
	}
	// send serialized to service for other operations
	resp, err := services.PermissionsService.FindAll(limit, offset)
	if err != nil {
		return c.JSON(err.Status(), err)
	}
	return c.JSON(resp.Status(), resp)
}

func (p permissionsController) Find(c echo.Context) error {
	pid, err := utils.ValidateAndCastToInt(c.Param("id"))
	if err != nil {
		return c.JSON(err.Status(), err)
	}
	resp, err := services.PermissionsService.Find(pid)
	if err != nil {
		return c.JSON(err.Status(), err)
	}
	return c.JSON(resp.Status(), resp)
}

func (p permissionsController) Update(c echo.Context) error {
	pid, err := utils.ValidateAndCastToInt(c.Param("id"))
	if err != nil {
		return c.JSON(err.Status(), err)
	}
	permission := new(domains.CreatePermissionsRequest)
	err = utils.ValidateAndBind(c, permission)
	if err != nil {
		return c.JSON(err.Status(), err)
	}
	resp, err := services.PermissionsService.Update(pid, *permission)
	if err != nil {
		return c.JSON(err.Status(), err)
	}
	return c.JSON(resp.Status(), resp)
}

func (p permissionsController) Delete(c echo.Context) error {
	pid, err := utils.ValidateAndCastToInt(c.Param("id"))
	if err != nil {
		return c.JSON(err.Status(), err)
	}
	resp, err := services.PermissionsService.Delete(pid)
	if err != nil {
		return c.JSON(err.Status(), err)
	}
	return c.JSON(resp.Status(), resp)
}
