package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func AuthLogin(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success login",
	})
}

func AuthRegister(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success register",
	})
}
