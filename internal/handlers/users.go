package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// PostUser godoc
// @Summary Create a new user
// @Description Create a new user
// @Tags users
// @Accept json
// @Produce json
// @Param user body User true "User details"
// @Success 200 {object} map[string]interface{}
// @Router /users [post]
func PostUser(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Post User Success",
	})
}

func GetUser(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Get User Success",
	})
}
