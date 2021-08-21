package utils

import (
	"github.com/alidevjimmy/go-rest-utils/rest_response"
	"github.com/bitokss/bitok-user-service/constants"
	"github.com/labstack/echo/v4"
	"strconv"
)

func ValidateAndCastToInt(param string) (int, rest_response.RestResp) {
	if param == "" {
		restErr := rest_response.NewBadRequestError(constants.InvalidInputErr, nil)
		return 0, restErr
	}
	castedParam, err := strconv.Atoi(param)
	if err != nil {
		restErr := rest_response.NewBadRequestError(constants.InvalidInputErr, nil)
		return 0, restErr
	}
	return castedParam, nil
}

func ValidateAndBind(c echo.Context , i interface{}) rest_response.RestResp{
	if err := c.Bind(i); err != nil {
		return rest_response.NewBadRequestError(constants.InvalidInputErr, nil)
	}
	if err := c.Validate(i); err != nil {
		return rest_response.NewBadRequestError(constants.InvalidInputErr, nil)
	}
	return nil
}