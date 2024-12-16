package main

import (
	"log"
	routeApp "meetapp/internal/handlers"
	utils "meetapp/pkg/database"
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
	db, err := utils.ConnectDB("saptoprasojo", "postgres", "meetappdb")
	if err != nil {
		log.Fatal("Failed to connect to database", err)
	}
	defer db.Close()

	route := echo.New()
	swag.ReadDoc()

	route.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	// reservation
	route.GET("/reservations", routeApp.GetAll)
	route.GET("/reservations/:id", routeApp.GetByID)
	route.PUT("/reservations/:id", routeApp.PutReservation)
	route.POST("/reservations", routeApp.PostReservation)

	// user
	route.GET("/users", routeApp.PostUser)

	//auth
	route.POST("/login", routeApp.AuthLogin)
	route.POST("/register", routeApp.AuthRegister)

	route.GET("/swagger/*", echoSwagger.WrapHandler)
	route.Logger.Fatal(route.Start(":7000"))
}
