package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetDasboard(c echo.Context) error {

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Dashboard",
	})
}
