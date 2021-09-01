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

func (l *levelsController) Create(c echo.Context) error {
	level := new(domains.CreateLevelsRequest)
	// validate level request data and bind it on CreateLevelsRequest struct
	err := utils.ValidateAndBind(c, level)
	if err != nil {
		return c.JSON(err.Status(), err)
	}
	resp, err := services.LevelsService.Create(*level)
	if err != nil {
		return c.JSON(err.Status(), err)
	}
	return c.JSON(resp.Status(), resp)
}

func (l *levelsController) FindAll(c echo.Context) error {
	limit := 50
	offset := 0
	limitParam := c.QueryParam("limit")
	offsetParam := c.QueryParam("offset")
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
	resp, err := services.LevelsService.FindAll(limit, offset)
	if err != nil {
		return c.JSON(err.Status(), err)
	}
	return c.JSON(resp.Status(), resp)
}

func (l *levelsController) Find(c echo.Context) error {
	lid, err := utils.ValidateAndCastToInt(c.Param("id"))
	if err != nil {
		return c.JSON(err.Status(), err)
	}
	resp, err := services.LevelsService.Find(lid)
	if err != nil {
		return c.JSON(err.Status(), err)
	}
	return c.JSON(resp.Status(), resp)
}

func (l *levelsController) Update(c echo.Context) error {
	lid, err := utils.ValidateAndCastToInt(c.Param("id"))
	if err != nil {
		return c.JSON(err.Status(), err)
	}
	level := new(domains.CreateLevelsRequest)
	err = utils.ValidateAndBind(c, level)
	if err != nil {
		return c.JSON(err.Status(), err)
	}
	resp, err := services.LevelsService.Update(uint(lid), *level)
	if err != nil {
		return c.JSON(err.Status(), err)
	}
	return c.JSON(resp.Status(), resp)
}

func (l *levelsController) Delete(c echo.Context) error {
	lid, err := utils.ValidateAndCastToInt(c.Param("id"))
	if err != nil {
		return c.JSON(err.Status(), err)
	}
	resp, err := services.LevelsService.Delete(lid)
	if err != nil {
		return c.JSON(err.Status(), err)
	}
	return c.JSON(resp.Status(), resp)
}
