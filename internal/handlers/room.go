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
// @Description Retrieve all rooms in the system. Requires a valid Bearer token.
// @Tags rooms
// @Produce json
// @Param Authorization header string true "Bearer <token>"
// @Success 200 {object} responseRoom
// @Failure 401 {object} map[string]string "Unauthorized or invalid token"
// @Failure 500 {object} map[string]string "Internal server error"
// @Security BearerAuth
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

// GetRoomByID godoc
// @Summary Get a room by ID
// @Description Get a room by ID
// @Tags rooms
// @Produce json
// @Param id path string true "Room ID"
// @Success 200 {object} responseRoom
// @Router /rooms/{id} [get]
func GetRoomByID(c echo.Context) error {
	id := c.Param("id")
	query := "SELECT id, name, type, capacity, price FROM rooms WHERE id = $1"
	var room Room

	err := utils.DB.QueryRow(query, id).Scan(&room.ID, &room.Name, &room.Type, &room.Capacity, &room.Price)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to retrieve room",
		})
	}

	response := responseRoom{
		Message: "Success",
		Data:    []Room{room},
	}
	return c.JSON(http.StatusOK, response)
}

// UpdateRoomByID godoc
// @Summary Update a room by ID
// @Description Update a room by ID
// @Tags rooms
// @Accept json
// @Produce json
// @Param id path string true "Room ID"
// @Param room body Room true "Room details"
// @Success 200 {object} responseRoom
// @Router /rooms/{id} [put]
func UpdateRoomByID(c echo.Context) error {
	id := c.Param("id")
	var room Room

	if err := c.Bind(&room); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request payload"})
	}

	// Update data di database
	query := "UPDATE rooms SET name = $1, type = $2, capacity = $3, price = $4 WHERE id = $5"
	result, err := utils.DB.Exec(query, room.Name, room.Type, room.Capacity, room.Price, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to update room",
		})
	}

	// Periksa apakah ada baris yang diperbarui
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Error checking affected rows"})
	}
	if rowsAffected == 0 {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "Room not found"})
	}

	// Ambil data terbaru setelah update
	selectQuery := "SELECT id, name, type, capacity, price FROM rooms WHERE id = $1"
	err = utils.DB.QueryRow(selectQuery, id).Scan(&room.ID, &room.Name, &room.Type, &room.Capacity, &room.Price)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to retrieve updated room"})
	}

	response := map[string]interface{}{
		"message": "Room updated successfully",
		"data":    room,
	}
	return c.JSON(http.StatusOK, response)
}

// DeleteRoomByID godoc
// @Summary Delete a room by ID
// @Description Delete a room by ID
// @Tags rooms
// @Produce json
// @Param id path string true "Room ID"
// Success 200 {object} map[string]string
// @Router /rooms/{id} [delete]
func DeleteRoomByID(c echo.Context) error {
	id := c.Param("id")
	query := "DELETE FROM rooms WHERE id = $1"
	_, err := utils.DB.Exec(query, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to retrieve user",
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Room deleted successfully",
	})
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
