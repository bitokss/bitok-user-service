package middlewares

import "github.com/labstack/echo/v4"

func OnlyWithPermissions(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(500, "you can't haji!")
	}
}
