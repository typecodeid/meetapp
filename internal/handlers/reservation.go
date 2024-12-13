package reservation

import (
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

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type Snack struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Category string `json:"category"`
	Package  string `json:"package"`
	Price    string `json:"price"`
}

type Reservation struct {
	ID           string `json:"id"`
	StartTime    string `json:"start_time"`
	EndTime      string `json:"end_time"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
	Status       string `json:"status"`
	Participants int    `json:"participants"`
	Name         string `json:"name"`
	Total_Snack  int    `json:"total_snack"`
	Company      string `json:"company"`
	Phone        string `json:"phone"`
	Room_price   int    `json:"room_price"`
	Final_price  int    `json:"final_price"`
	User         User   `json:"user"`
	Room         Room   `json:"room"`
	Snack        Snack  `json:"snack"`
}

var data = []Reservation{
	{
		ID:           "1",
		StartTime:    "2022-01-01T10:00:00Z",
		EndTime:      "2022-01-01T11:00:00Z",
		CreatedAt:    "2022-01-01T10:00:00Z",
		UpdatedAt:    "2022-01-01T10:00:00Z",
		Status:       "Booked",
		Participants: 7,
		Name:         "Monkey D Luffy",
		Total_Snack:  7,
		Company:      "One Piece",
		Phone:        "1234567890",
		Room_price:   700000,
		Final_price:  800000,
		User: User{
			ID:       "1",
			Username: "John Doe",
			Email:    "5o0rI@example.com",
		},
		Room: Room{
			ID:       "1",
			Name:     "Room 1",
			Type:     "Meeting Room",
			Capacity: 10,
			Price:    100,
		},
		Snack: Snack{
			ID:       "1",
			Name:     "Snack 1",
			Category: "Food",
			Package:  "Small",
			Price:    "10",
		},
	},
	{
		ID:           "2",
		StartTime:    "2022-01-01T10:00:00Z",
		EndTime:      "2022-01-01T11:00:00Z",
		CreatedAt:    "2022-01-01T10:00:00Z",
		UpdatedAt:    "2022-01-01T10:00:00Z",
		Status:       "Paid",
		Participants: 10,
		Name:         "Sun Goku",
		Total_Snack:  7,
		Company:      "Dragon Ball",
		Phone:        "1234567890",
		Room_price:   800000,
		Final_price:  900000,
		User: User{
			ID:       "1",
			Username: "John Doe",
			Email:    "5o0rI@example.com",
		},
		Room: Room{
			ID:       "1",
			Name:     "Room 2",
			Type:     "Meeting Room",
			Capacity: 10,
			Price:    100,
		},
		Snack: Snack{
			ID:       "1",
			Name:     "Snack 1",
			Category: "Food",
			Package:  "Small",
			Price:    "10",
		},
	},
}

// GetAll godoc
// @Summary Get all reservations
// @Description Retrieve all reservations in the system
// @Tags reservations
// @Accept json
// @Produce json
// @Success 200 {array} Reservation
// @Router /reservations [get]
func GetAll(c echo.Context) error {
	type Response struct {
		Message string        `json:"message"`
		Data    []Reservation `json:"data"`
	}

	response := Response{
		Message: "Success",
		Data:    data,
	}
	return c.JSON(http.StatusOK, response)
}

// GetByID godoc
// @Summary Get a reservation by ID
// @Description Retrieve a reservation by its ID
// @Tags reservations
// @Accept json
// @Produce json
// @Param id path string true "Reservation ID"
// @Success 200 {object} Reservation
// @Failure 404 {object} map[string]string
// @Router /reservations/{id} [get]
func GetByID(c echo.Context) error {
	id := c.Param("id")
	for _, reservation := range data {
		if reservation.ID == id {
			return c.JSON(http.StatusOK, map[string]interface{}{
				"message": "Success",
				"data":    reservation,
			})
		}
	}

	// Jika data tidak ditemukan
	return c.JSON(http.StatusNotFound, map[string]string{
		"message": "Reservation not found",
	})
}

// PostReservation godoc
// @Summary Create a new reservation
// @Description Create a new reservation
// @Tags reservations
// @Accept json
// @Produce json
// @Param reservation body Reservation true "Reservation details"
// @Success 200 {object} map[string]string
// @Router /reservations [post]
func PostReservation(c echo.Context) error {
	var reservation Reservation
	if err := c.Bind(&reservation); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Invalid request body",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
		"data":    reservation,
	})
}
