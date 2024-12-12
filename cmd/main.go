package main

import (
	reservation "meetapp/internal/handlers"
	"net/http"

	_ "meetapp/docs"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger" // echo-swagger middleware
	"github.com/swaggo/swag"
)

// @title Swagger MeetApp By Sinau Koding API
// @version 1.0
// @description This is docomentation api from swagger
// @termsOfService http://swagger.io/terms/

// GetAll godoc
// @Summary Get all reservations
// @Description Retrieve all reservations in the system
// @Tags reservations
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /reservations [get]

func main() {
	swag.ReadDoc()
	route := echo.New()

	route.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	route.GET("/swagger/*", echoSwagger.WrapHandler)

	route.GET("/reservations", reservation.GetAll)
	route.GET("/reservations/:id", reservation.GetByID)
	route.Logger.Fatal(route.Start(":7000"))
}
