package controllers

import (
	"github.com/alidevjimmy/go-rest-utils/rest_response"
	"github.com/bitokss/bitok-user-service/constants"
	"github.com/bitokss/bitok-user-service/domains/v1"
	"github.com/bitokss/bitok-user-service/services/v1"
	"github.com/bitokss/bitok-user-service/utils"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

var (
	UsersController usersControllerInterface = &usersController{}
)

type usersControllerInterface interface {
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

type usersController struct{}

func (u *usersController) Create(c echo.Context) error {
	user := new(domains.CreateUsersRequest)
	// validate user request data and bind it on CreateUsersRequest struct
	err := utils.ValidateAndBind(c, user)
	if err != nil {
		return c.JSON(err.Status(), err)
	}
	resp, err := services.UsersService.Create(*user)
	if err != nil {
		return c.JSON(err.Status(), err)
	}
	return c.JSON(resp.Status(), resp)
}

func (u *usersController) FindAll(c echo.Context) error {
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
	resp, err := services.UsersService.FindAll(limit, offset)
	if err != nil {
		return c.JSON(err.Status(), err)
	}
	return c.JSON(resp.Status(), resp)
}

func (u *usersController) Find(c echo.Context) error {
	uid, err := utils.ValidateAndCastToInt(c.Param("id"))
	if err != nil {
		return c.JSON(err.Status(), err)
	}
	resp, err := services.UsersService.Find(uid)
	if err != nil {
		return c.JSON(err.Status(), err)
	}
	return c.JSON(resp.Status(), resp)
}

func (u *usersController) Update(c echo.Context) error {
	uid, err := utils.ValidateAndCastToInt(c.Param("id"))
	if err != nil {
		return c.JSON(err.Status(), err)
	}
	user := new(domains.CreateUsersRequest)
	err = utils.ValidateAndBind(c, user)
	if err != nil {
		return c.JSON(err.Status(), err)
	}
	resp, err := services.UsersService.Update(uid, *user)
	if err != nil {
		return c.JSON(err.Status(), err)
	}
	return c.JSON(resp.Status(), resp)
}

func (u *usersController) Delete(c echo.Context) error {
	uid, err := utils.ValidateAndCastToInt(c.Param("id"))
	if err != nil {
		return c.JSON(err.Status(), err)
	}
	resp, err := services.UsersService.Delete(uid)
	if err != nil {
		return c.JSON(err.Status(), err)
	}
	return c.JSON(resp.Status(), resp)
}
func (u *usersController) Register(c echo.Context) error {
	panic("implement me")
}

func (u *usersController) Login(c echo.Context) error {
	panic("implement me")
}

func (u *usersController) FindByToken(c echo.Context) error {
	token := c.Param("token")
	if token == "" {
		return c.JSON(http.StatusBadRequest, rest_response.NewBadRequestError(constants.InvalidInputErr , nil))
	}
	resp, err := services.UsersService.FindByToken(token)
	if err != nil {
		return c.JSON(err.Status(), err)
	}
	return c.JSON(resp.Status(), resp)
}

func (u *usersController) FindByUsername(c echo.Context) error {
	panic("implement me")
}

func (u *usersController) ResetPassword(c echo.Context) error {
	panic("implement me")
}

func (u *usersController) TickRequest(c echo.Context) error {
	panic("implement me")
}