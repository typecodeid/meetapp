package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// GetSnack godoc
// @Summary Get Snack
// @Description Get Snack
// @Tags Snack
// @Produce json
// @Success 200 {object} map[string]string
// @Router /snack [get]
func GetSnack(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Get Snack Here",
	})
}

// CreateSnack godoc
// @Summary Create Snack
// @Description Create Snack
// @Tags Snack
// @Produce json
// @Success 200 {object} map[string]string
// @Router /snack [post]
func CreateSnack(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Create Snack Here",
	})
}
