package main

import (
	reservation "meetapp/internal/handlers"
	"net/http"

	_ "meetapp/docs" // Pastikan impor ini ada

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger" // echo-swagger middleware
	"github.com/swaggo/swag"
)

// @title Swagger MeetApp By Sinau Koding API
// @version 1.0
// @description This is documentation API from Swagger
// @termsOfService http://swagger.io/terms/
// @host localhost:7000
// @BasePath /

func main() {
	route := echo.New()
	swag.ReadDoc()

	route.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	route.GET("/reservations", reservation.GetAll)
	route.GET("/reservations/:id", reservation.GetByID)
	route.GET("/swagger/*", echoSwagger.WrapHandler)
	route.Logger.Fatal(route.Start(":7000"))
}
