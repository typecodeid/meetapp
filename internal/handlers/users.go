package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

var users = make(map[string]User)
var rooms = make(map[string]Room)
var reservation = make(map[string]Reservation)
var snacks = make(map[string]Snack)

var Users = []User{
	{ID: 1, Username: "John Doe", Gmail: "Johndoe@gmail.com", image: "image", password: "123", role: "admin", status: "status", language: "bahasa"},
	{ID: 2, Username: "Sun Goku", Gmail: "Sungoku@gmail.com", image: "image", password: "456", role: "staff", status: "status", language: "bahasa"},
	{ID: 3, Username: "Sun Goku", Gmail: "Sungoku2@gmail.com", image: "image", password: "789", role: "manager", status: "status", language: "bahasa"},
}

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

func GetUserByID(c echo.Context) error {
	id := c.Param("id")
	user, exist := users[id]
	if !exist {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "user not found"})
	}
	return c.JSON(http.StatusOK, user)
}

func CreateUser(c echo.Context) error {
	var user User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid input"})
	}
	users[user.ID] = user
	return c.JSON(http.StatusCreated, user)
}

func GetUser(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Get User Success",
	})
}

// handler room
func GetRoom(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Get Room Success",
	})
}

var Room = []Room{
	{ID: 1, Name: "Nilna", Type: "small", Capacity: 10, Price: 100},
	{ID: 2, Name: "Alminah", Type: "medium", Capacity: 10, Price: 100},
	{ID: 1, Name: "Nilna", Type: "large", Capacity: 10, Price: 100},
}

func CreateRoom(c echo.Context) error {
	var room Room
	if err := c.Bind(&room); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid input"})
	}
	rooms[room.ID] = room
	return c.JSON(http.StatusCreated, room)

}

// handler untuk reservation
func CreateReservationForRoom(c echo.Context) error {
	RoomID := c.Param("id")
	var res Reservation
	if err := c.Bind(&res); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid input"})
	}
	res.Room.ID = RoomID
	reservation[res.ID] = res
	return c.JSON(http.StatusCreated, res)
}

// handler untuk snack
func GetSnack(c echo.Context) error {
	return c.JSON(http.StatusOK, snacks)
}

var Snack = []Snack{
	{ID: 1, Name: "Snack 1", Category: "Food", Package: "Small", Price: "10"},
	{ID: 2, Name: "Snack 1", Category: "Food", Package: "Small", Price: "10"},
	{ID: 2, Name: "Snack 1", Category: "Food", Package: "Small", Price: "10"},
}

func CreateSnack(c echo.Context) error {
	var snack Snack
	if err := c.Bind(&snack); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid input"})
	}
	snacks[snack.ID] = snack
	return c.JSON(http.StatusCreated, snack)
}
