package handlers

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type Rooms struct {
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
	ImageID  string `json:"image_id"`
	Password string `json:"password"`
	Role     string `json:"role"`
	Status   bool   `json:"status"`
	Language string `json:"language"`
}

type Snacks struct {
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
	Room         Rooms  `json:"room"`
	Snack        Snacks `json:"snack"`
}

var data = []Reservation{
	{
		ID:           "1",
		StartTime:    "2022-01-01T10:00:00Z",
		EndTime:      "2022-01-01T11:00:00Z",
		CreatedAt:    "2022-01-01T10:00:00Z",
		UpdatedAt:    "2022-01-01T10:00:00Z",
		Status:       "booked",
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
		Room: Rooms{
			ID:       "1",
			Name:     "Ruang 1",
			Type:     "small",
			Capacity: 10,
			Price:    100,
		},
		Snack: Snacks{
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
		Status:       "paid",
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
		Room: Rooms{
			ID:       "2",
			Name:     "Ruang 2",
			Type:     "medium",
			Capacity: 10,
			Price:    100,
		},
		Snack: Snacks{
			ID:       "1",
			Name:     "Snack 1",
			Category: "Food",
			Package:  "Small",
			Price:    "10",
		},
	},
	{
		ID:           "3",
		StartTime:    "2022-01-01T10:00:00Z",
		EndTime:      "2022-01-01T11:00:00Z",
		CreatedAt:    "2022-01-01T10:00:00Z",
		UpdatedAt:    "2022-01-01T10:00:00Z",
		Status:       "cancel",
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
		Room: Rooms{
			ID:       "3",
			Name:     "Room 3",
			Type:     "large",
			Capacity: 10,
			Price:    100,
		},
		Snack: Snacks{
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
// @Param status query string false "Filter by status"
// @Param room_type query string false "Filter by room type"
// @Param start_date query string false "Filter by start date (YYYY-MM-DD)"
// @Param end_date query string false "Filter by end date (YYYY-MM-DD)"
// @Success 200 {array} Reservation
// @Router /reservations [get]
func GetAll(c echo.Context) error {
	type Response struct {
		Message string        `json:"message"`
		Data    []Reservation `json:"data"`
	}

	status := c.QueryParam("status")
	roomType := c.QueryParam("room_type")
	startDateStr := c.QueryParam("start_date") // rename dari starDate ke startDateStr
	endDateStr := c.QueryParam("end_date")     // rename dari endDate ke endDateStr

	var startDate, endDate time.Time
	var err error

	// Parse start_date jika ada
	if startDateStr != "" {
		startDate, err = time.Parse("2006-01-02", startDateStr)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"message": "Invalid start date format",
			})
		}
	}

	// Parse end_date jika ada
	if endDateStr != "" {
		endDate, err = time.Parse("2006-01-02", endDateStr)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"message": "Invalid end date format",
			})
		}
	}

	filteredData := []Reservation{}

	for _, r := range data {
		// Filter status
		if status != "" && r.Status != status {
			continue
		}

		// Filter room_type
		if roomType != "" && r.Room.Type != roomType {
			continue
		}

		// Parsing StartTime & EndTime dari reservation
		reservationStart, errStart := time.Parse(time.RFC3339, r.StartTime)
		reservationEnd, errEnd := time.Parse(time.RFC3339, r.EndTime)
		if errStart != nil || errEnd != nil {
			// Jika data tidak valid, skip saja
			continue
		}

		// Filter berdasarkan start_date
		if !startDate.IsZero() && reservationStart.Before(startDate) {
			continue
		}

		// Filter berdasarkan end_date
		if !endDate.IsZero() && reservationEnd.After(endDate) {
			continue
		}

		// Jika semua filter lolos, masukkan ke filteredData
		filteredData = append(filteredData, r)
	}

	response := Response{
		Message: "Success",
		Data:    filteredData, // Pastikan menggunakan filteredData, bukan data
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

// PutReservation godoc
// @Summary Edit a reservation
// @Description Edit a reservation
// @Tags reservations
// @Accept json
// @Produce json
// @Param id path string true "Reservation ID"
// @Success 200 {object} map[string]string
// @Router /reservations/{id} [put]
func PutReservation(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Edit data",
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
