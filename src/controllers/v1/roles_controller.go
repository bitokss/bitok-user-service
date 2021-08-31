package controllers

import (
	"strconv"

	"github.com/alidevjimmy/go-rest-utils/rest_response"
	"github.com/bitokss/bitok-user-service/constants"
	"github.com/bitokss/bitok-user-service/domains/v1"
	"github.com/bitokss/bitok-user-service/services/v1"
	"github.com/bitokss/bitok-user-service/utils"
	"github.com/labstack/echo/v4"
)

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

func (r *rolesController) Create(c echo.Context) error {
	role := new(domains.CreateRolesRequest)
	// validate role request data and bind it on CreateRolesRequest struct
	err := utils.ValidateAndBind(c, role)
	if err != nil {
		return c.JSON(err.Status(), err)
	}
	resp, err := services.RolesService.Create(*role)
	if err != nil {
		return c.JSON(err.Status(), err)
	}
	return c.JSON(resp.Status(), resp)
}

func (r *rolesController) FindAll(c echo.Context) error {
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
	resp, err := services.RolesService.FindAll(limit, offset)
	if err != nil {
		return c.JSON(err.Status(), err)
	}
	return c.JSON(resp.Status(), resp)
}

func (r *rolesController) Find(c echo.Context) error {
	rid, err := utils.ValidateAndCastToInt(c.Param("id"))
	if err != nil {
		return c.JSON(err.Status(), err)
	}
	resp, err := services.RolesService.Find(rid)
	if err != nil {
		return c.JSON(err.Status(), err)
	}
	return c.JSON(resp.Status(), resp)
}

func (r *rolesController) Update(c echo.Context) error {
	rid, err := utils.ValidateAndCastToInt(c.Param("id"))
	if err != nil {
		return c.JSON(err.Status(), err)
	}
	role := new(domains.CreateRolesRequest)
	err = utils.ValidateAndBind(c, role)
	if err != nil {
		return c.JSON(err.Status(), err)
	}
	resp, err := services.RolesService.Update(uint(rid), *role)
	if err != nil {
		return c.JSON(err.Status(), err)
	}
	return c.JSON(resp.Status(), resp)
}

func (r *rolesController) Delete(c echo.Context) error {
	rid, err := utils.ValidateAndCastToInt(c.Param("id"))
	if err != nil {
		return c.JSON(err.Status(), err)
	}
	resp, err := services.RolesService.Delete(rid)
	if err != nil {
		return c.JSON(err.Status(), err)
	}
	return c.JSON(resp.Status(), resp)
}
