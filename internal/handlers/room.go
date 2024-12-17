package handlers

import (
	utils "meetapp/pkg/database"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Room struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Type     string `json:"type"`
	Capacity int    `json:"capacity"`
	Price    int    `json:"price"`
}

type responseRoom struct {
	Message string `json:"message"`
	Data    []Room `json:"data"`
}

// GetRooms godoc
// @Summary Get all rooms
// @Description Retrieve all rooms in the system
// @Tags rooms
// @Produce json
// @Success 200 {object} responseRoom
// @Router /rooms [get]
func GetRooms(c echo.Context) error {
	query := "SELECT id, name, type, capacity, price FROM rooms"

	rows, err := utils.DB.Query(query)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to retrieve users",
		})
	}
	var room []Room

	for rows.Next() {
		var roomData Room
		err := rows.Scan(&roomData.ID, &roomData.Name, &roomData.Type, &roomData.Capacity, &roomData.Price)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": "Failed to scan user data",
			})
		}
		room = append(room, roomData)
	}

	response := responseRoom{
		Message: "Success",
		Data:    room,
	}
	return c.JSON(http.StatusOK, response)
}

// CreateRoom godoc
// @Summary Create a new room
// @Description Create a new room
// @Tags rooms
// @Accept json
// @Produce json
// @Param room body Room true "Room details"
// @Success 201 {object} Room
// @Router /rooms [post]
func CreateRoom(c echo.Context) error {
	var input Room

	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request payload"})
	}

	// cek nama ruangan apakah sudah ada?
	var exists bool
	checkquery := "SELECT EXISTS (SELECT 1 FROM rooms WHERE name = $1)"
	err := utils.DB.QueryRow(checkquery, input.Name).Scan(&exists)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Error checking room name"})
	}
	if exists {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Room name already exists"})
	}

	// Simpan data ke database
	query := "INSERT INTO rooms (name, type, capacity, price) VALUES ($1, $2, $3, $4) RETURNING id"
	var id string
	errInsert := utils.DB.QueryRow(query, input.Name, input.Type, input.Capacity, input.Price).Scan(&id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": errInsert.Error()})
	}
	input.ID = id

	return c.JSON(http.StatusCreated, input)
}
