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
	user := "postgres"
	password := "PmdkTSGxDXYqVFKfwwmkOYgMftbwwdXs"
	dbname := "railway"
	host := "junction.proxy.rlwy.net"
	port := "35001"

	db, err := utils.ConnectDB(user, password, dbname, host, port)
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
	route.GET("/reservations", routeApp.GetAllReservation)
	route.GET("/reservations/:id", routeApp.GetByID)
	route.PUT("/reservations/:id", routeApp.PutReservation)
	route.POST("/reservations", routeApp.PostReservation)

	// user
	route.GET("/users", routeApp.GetUsers)
	route.GET("/users/:id", routeApp.GetUserByID)
	route.PUT("/users/:id", routeApp.UpdateUserByID)
	route.DELETE("/users/:id", routeApp.DeleteUserByID)

	// room
	route.GET("/rooms", routeApp.GetRooms)
	route.GET("/rooms/:id", routeApp.GetRoomByID)
	route.POST("/rooms", routeApp.CreateRoom)
	route.PUT("/rooms/:id", routeApp.UpdateRoomByID)
	route.DELETE("/rooms/:id", routeApp.DeleteRoomByID)
	// route.POST("/room/:id/reservation", routeApp.CreateReservationForRoom)

	// snack
	route.GET("/snack", routeApp.GetSnack)
	route.POST("/snack", routeApp.CreateSnack)

	//auth
	route.POST("/login", routeApp.AuthLogin)
	route.POST("/register", routeApp.AuthRegister)

	route.GET("/swagger/*", echoSwagger.WrapHandler)
	route.Logger.Fatal(route.Start(":7000"))
}
