package middlewares

import (
	"github.com/alidevjimmy/go-rest-utils/rest_response"
	"github.com/bitokss/bitok-user-service/constants"
	"github.com/labstack/echo/v4"
	"net/http"
)

func OnlyWithPermissions(permission []string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			token := c.Request().Header.Get(echo.HeaderAuthorization)
			if token == "" {
				er := rest_response.NewUnauthorizedError(constants.UnAuthorizedErr , nil)
				return c.JSON(http.StatusUnauthorized, er)
			}
			// get user by token and check his permission
			return next(c)
		}
	}
}
