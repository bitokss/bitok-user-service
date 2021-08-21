package middlewares

import (
	"github.com/alidevjimmy/go-rest-utils/rest_response"
	"github.com/bitokss/bitok-user-service/constants"
	"github.com/bitokss/bitok-user-service/services/v1"
	"github.com/labstack/echo/v4"
	"net/http"
)

func OnlyLogin (next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Request().Header.Get(echo.HeaderAuthorization)
		if token == "" {
			return c.JSON(http.StatusUnauthorized, rest_response.NewUnauthorizedError(constants.UnAuthorizedErr , nil))
		}
		_ , err := services.UsersService.FindByToken(token)
		if err != nil {
			return c.JSON(err.Status(),err)
		}
		return next(c)
	}
}