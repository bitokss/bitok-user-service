package middlewares

import (
	"github.com/alidevjimmy/go-rest-utils/rest_response"
	"github.com/bitokss/bitok-user-service/constants"
	"github.com/bitokss/bitok-user-service/services/v1"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

func OnlyLogin (next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Request().Header.Get(echo.HeaderAuthorization)
		if token == "" {
			return c.JSON(http.StatusUnauthorized, rest_response.NewUnauthorizedError(constants.UnAuthorizedErr , nil))
		}
		splitToken := strings.Split(token, "Bearer ")
		if len(splitToken) != 2 {
			return c.JSON(http.StatusUnauthorized, rest_response.NewUnauthorizedError(constants.UnAuthorizedErr , nil))
		}
		token = splitToken[1]
		_ , err := services.UsersService.FindByToken(token)
		if err != nil {
			return c.JSON(err.Status(),err)
		}
		return next(c)
	}
}