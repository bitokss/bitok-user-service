package middlewares

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/alidevjimmy/go-rest-utils/rest_response"
	"github.com/bitokss/bitok-user-service/constants"
	"github.com/bitokss/bitok-user-service/domains/v1"
	"github.com/bitokss/bitok-user-service/services/v1"
	"github.com/labstack/echo/v4"
)

func OnlyOwner(next echo.HandlerFunc) echo.HandlerFunc {
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
		resp, err := services.UsersService.FindByToken(token)
		if err != nil {
			return c.JSON(err.Status(), err)
		}
		user, ok := resp.Data().(domains.UserResp)
		if !ok {
			return c.JSON(http.StatusInternalServerError, rest_response.NewInternalServerError(constants.InternalServerErr, nil))
		}
		userID, er := strconv.Atoi(c.Param("user_id"))
		if er != nil {
			return c.JSON(http.StatusBadRequest, rest_response.NewBadRequestError(constants.InvalidInputErr, nil))

		}
		if user.ID != uint(userID) {
			return c.JSON(http.StatusBadRequest, rest_response.NewBadRequestError(constants.UnAuthorizedErr, nil))
		}
		return next(c)
	}
}
