package main

import (
	reservation "meetapp/internal/handlers"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	route := echo.New()
	route.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	route.GET("/reservations", reservation.GetAll)
	route.GET("/reservations/:id", reservation.GetByID)
	route.Logger.Fatal(route.Start(":7000"))
}
