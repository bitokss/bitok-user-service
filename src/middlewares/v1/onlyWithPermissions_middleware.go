package middlewares

import (
	"github.com/alidevjimmy/go-rest-utils/rest_response"
	"github.com/bitokss/bitok-user-service/constants"
	"github.com/bitokss/bitok-user-service/domains/v1"
	"github.com/bitokss/bitok-user-service/services/v1"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

func OnlyWithPermissions(permissions []string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			token := c.Request().Header.Get(echo.HeaderAuthorization)
			if token == "" {
				return c.JSON(http.StatusUnauthorized, rest_response.NewUnauthorizedError(constants.UnAuthorizedErr, nil))
			}
			splitToken := strings.Split(token, "Bearer ")
			if len(splitToken) != 2 {
				return c.JSON(http.StatusUnauthorized, rest_response.NewUnauthorizedError(constants.UnAuthorizedErr, nil))
			}
			token = splitToken[1]
			if len(permissions) == 0 {
				return next(c)
			}
			resp, err := services.UsersService.FindByToken(token)
			if err != nil {
				return c.JSON(err.Status(), err)
			}
			user, ok := resp.Data().(domains.UserResp)
			if !ok {
				return c.JSON(http.StatusInternalServerError, rest_response.NewInternalServerError(constants.InternalServerErr, nil))
			}
			// O(n+m) checker for permissions
			checker := map[string]bool{}
			for _, v := range permissions {
				checker[v] = false
			}
			for _, r := range user.Roles {
				for _, p := range r.Permissions {
					checker[p.Symbol] = true
				}
			}
			for _, v := range checker {
				if v == false {
					return c.JSON(http.StatusUnauthorized, rest_response.NewUnauthorizedError(constants.UnAuthorizedErr, nil))
				}
			}
			return next(c)
		}
	}
}
