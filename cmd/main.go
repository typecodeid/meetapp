package main

import (
	"log"
	routeApp "meetapp/internal/handlers"
	"meetapp/internal/mymiddleware"
	utils "meetapp/pkg/database"
	"net/http"

	_ "meetapp/docs" // Pastikan impor ini ada

	"github.com/labstack/echo/v4"

	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger" // echo-swagger middleware
	"github.com/swaggo/swag"
)

// @title Swagger MeetApp By Sinau Koding API
// @version 1.0
// @description This is documentation API from Swagger. Jika ada masalah token silahkan coba pakai postman
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
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

	route.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"https://example.com", "http://localhost:3000"}, // Domain yang diizinkan
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
		AllowHeaders: []string{"Content-Type", "Authorization"},
	}))

	swag.ReadDoc()
	route.Static("/uploads", "uploads")
	route.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	// Dashboard
	route.GET("/dashboard", routeApp.GetDashboard, mymiddleware.TokenRole("admin"))

	// reservation
	route.GET("/reservations", routeApp.GetAllReservation, mymiddleware.TokenRole("user"))
	route.GET("/reservations/:id", routeApp.GetByID, mymiddleware.TokenRole("user"))
	route.PUT("/reservations/:id", routeApp.PutReservation, mymiddleware.TokenRole("user"))
	route.POST("/reservations", routeApp.PostReservation, mymiddleware.TokenRole("user"))

	// user
	route.GET("/users", routeApp.GetUsers, mymiddleware.TokenRole("admin"))
	route.GET("/users/:id", routeApp.GetUserByID, mymiddleware.TokenRole("user"))
	route.PUT("/users/:id", routeApp.UpdateUserByID, mymiddleware.TokenRole("admin"))
	route.DELETE("/users/:id", routeApp.DeleteUserByID, mymiddleware.TokenRole("admin"))

	// room
	route.GET("/rooms", routeApp.GetRooms, mymiddleware.TokenRole("user"))
	route.GET("/rooms/:id", routeApp.GetRoomByID, mymiddleware.TokenRole("user"))
	route.POST("/rooms", routeApp.CreateRoom, mymiddleware.TokenRole("admin"))
	route.PUT("/rooms/:id", routeApp.UpdateRoomByID, mymiddleware.TokenRole("admin"))
	route.DELETE("/rooms/:id", routeApp.DeleteRoomByID, mymiddleware.TokenRole("admin"))

	// snack
	route.GET("/snack", routeApp.GetSnack, mymiddleware.TokenRole("user"))

	//auth
	route.POST("/login", routeApp.AuthLogin)
	route.POST("/register", routeApp.AuthRegister)

	// images
	route.POST("/images", routeApp.PostImage, mymiddleware.TokenRole("user"))

	route.GET("/swagger/*", echoSwagger.WrapHandler)
	route.Logger.Fatal(route.Start(":7000"))
}
