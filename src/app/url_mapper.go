package app

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func urlMapper() {
	e.GET("/" , func(c echo.Context) error {
		return c.String(http.StatusOK,"Test!")
	})
}