package handlers

import (
	utils "meetapp/pkg/database"
	"net/http"

	"github.com/labstack/echo/v4"
)

type response struct {
	Message string         `json:"message"`
	Data    []UserResponse `json:"data"`
}

// handler user
// @summary Get all users
// @description Get all users
// @tags users
// @produce json
// @success 200 {object} response
// @router /users [get]
func GetUsers(c echo.Context) error {
	query := "SELECT id, image_id, username, email, role, status, language FROM users"

	rows, err := utils.DB.Query(query)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to retrieve users",
		})
	}

	var userResponses []UserResponse

	for rows.Next() {
		var userResponse UserResponse
		err := rows.Scan(&userResponse.ID, &userResponse.ImageID, &userResponse.Username, &userResponse.Email, &userResponse.Role, &userResponse.Status, &userResponse.Language)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": "Failed to scan user data",
			})
		}

		userResponses = append(userResponses, userResponse)
	}

	response := response{
		Message: "Success",
		Data:    userResponses,
	}
	return c.JSON(http.StatusOK, response)
}

// handler user
// @summary Get user by ID
// @description Get user by ID
// @tags users
// @produce json
// @Param id path string true "Reservation ID"
// @success 200 {object} response
// @router /users/{id} [get]
func GetUserByID(c echo.Context) error {
	query := "SELECT id, image_id, username, email, role, status, language FROM users WHERE id = $1"

	id := c.Param("id")
	var userResponse UserResponse

	err := utils.DB.QueryRow(query, id).Scan(&userResponse.ID, &userResponse.ImageID, &userResponse.Username, &userResponse.Email, &userResponse.Role, &userResponse.Status, &userResponse.Language)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to retrieve user",
		})
	}

	response := response{
		Message: "Success",
		Data:    []UserResponse{userResponse},
	}
	return c.JSON(http.StatusOK, response)
}

// func CreateReservationForRoom(c echo.Context) error {
// 	RoomID := c.Param("id")
// 	var res Reservation
// 	if err := c.Bind(&res); err != nil {
// 		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid input"})
// 	}
// 	res.Room.ID = RoomID
// 	reservation[res.ID] = res
// 	return c.JSON(http.StatusCreated, res)
// }

// @summary Update user by ID
// @description Update user by ID
// @tags users
// @produce json
// @Param id path string true "Reservation ID"
// @success 200 {object} response
// @router /users/{id} [put]
func UpdateUserByID(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Update User Here",
	})
}

// @summary Delete user by ID
// @description Delete user by ID
// @tags users
// @produce json
// @Param id path string true "Reservation ID"
// @success 200 {object} response
// @router /users/{id} [delete]
func DeleteUserByID(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Delete User Here",
	})
}
