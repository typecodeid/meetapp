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
// @Security BearerAuth
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
// @Security BearerAuth
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
// @Security BearerAuth
// @router /users/{id} [put]
func UpdateUserByID(c echo.Context) error {
	id := c.Param("id")
	var user User

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusOK, map[string]string{"message": "Invalid request payload"})
	}

	//update data di database
	query := "UPDATE user SET name = $1, image_id =$2, username = $3, email = $4, role = $5, status = $6, Language = $7 WHERE id = $8 "
	result, err := utils.DB.Exec(query, user.ImageID, user.Username, user.Email, user.Role, user.Status, user.Language, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to update user",
		})
	}

	// Periksa apakah ada baris yang diperbarui
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Error checking affected rows"})
	}
	if rowsAffected == 0 {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "User not found"})
	}

	// Ambil data terbaru setelah update
	selectQuery := "SELECT id, image_id, username, email, role, status, language  FROM user WHERE id = $1"
	err = utils.DB.QueryRow(selectQuery, id).Scan(&user.ID, &user.ImageID, &user.Username, &user.Email, &user.Role, &user.Status, &user.Language)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to retrieve updated room"})
	}

	response := map[string]interface{}{
		"message": "User updated successfully",
		"data":    user,
	}
	return c.JSON(http.StatusOK, response)
}

// @summary Delete user by ID
// @description Delete user by ID
// @tags users
// @produce json
// @Param id path string true "Reservation ID"
// @success 200 {object} response
// @Security BearerAuth
// @router /users/{id} [delete]
func DeleteUserByID(c echo.Context) error {
	id := c.Param("id")
	query := "DELETE FROM user WHERE id = $1"
	_, err := utils.DB.Exec(query, id)
	if err != nil {
		return c.JSON(http.StatusOK, map[string]string{
			"message": "Failed to delete user",
		})
	}
	return c.JSON(http.StatusOK, map[string]string{
		"message": "user deleted succesfully",
	})
}
