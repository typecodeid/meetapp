package main

import (
	"log"
	routeApp "meetapp/internal/handlers"
	"meetapp/internal/middleware"
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

	// Dashboard
	route.GET("/dashboard", routeApp.GetDasboard, middleware.TokenRole("admin"))

	// reservation
	route.GET("/reservations", routeApp.GetAllReservation, middleware.TokenRole("user"))
	route.GET("/reservations/:id", routeApp.GetByID, middleware.TokenRole("user"))
	route.PUT("/reservations/:id", routeApp.PutReservation, middleware.TokenRole("user"))
	route.POST("/reservations", routeApp.PostReservation, middleware.TokenRole("user"))

	// user
	route.GET("/users", routeApp.GetUsers, middleware.TokenRole("user"))
	route.GET("/users/:id", routeApp.GetUserByID, middleware.TokenRole("user"))
	route.PUT("/users/:id", routeApp.UpdateUserByID, middleware.TokenRole("user"))
	route.DELETE("/users/:id", routeApp.DeleteUserByID, middleware.TokenRole("admin"))

	// room
	route.GET("/rooms", routeApp.GetRooms, middleware.TokenRole("user"))
	route.GET("/rooms/:id", routeApp.GetRoomByID, middleware.TokenRole("user"))
	route.POST("/rooms", routeApp.CreateRoom, middleware.TokenRole("admin"))
	route.PUT("/rooms/:id", routeApp.UpdateRoomByID, middleware.TokenRole("admin"))
	route.DELETE("/rooms/:id", routeApp.DeleteRoomByID, middleware.TokenRole("admin"))

	// snack
	route.GET("/snack", routeApp.GetSnack, middleware.TokenRole("user"))
	route.POST("/snack", routeApp.CreateSnack, middleware.TokenRole("user"))

	//auth
	route.POST("/login", routeApp.AuthLogin)
	route.POST("/register", routeApp.AuthRegister)

	route.GET("/swagger/*", echoSwagger.WrapHandler)
	route.Logger.Fatal(route.Start(":7000"))
}
