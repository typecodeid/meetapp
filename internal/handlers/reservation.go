package handlers

import (
	"fmt"
	utils "meetapp/pkg/database"
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
	ID                string `json:"id"`
	RoomID            string `json:"room_id"`
	UserID            string `json:"user_id"`
	SnackID           string `json:"snack_id"`
	StartTime         string `json:"start_time"`
	EndTime           string `json:"end_time"`
	BookingDate       string `json:"booking_date"`
	CreatedAt         string `json:"created_at"`
	UpdatedAt         string `json:"updated_at"`
	Status            string `json:"status"`
	Participants      int    `json:"participants"`
	Name              string `json:"name"`
	Total_Snack       int    `json:"total_snack"`
	Total_Snack_Price int    `json:"total_snack_price"`
	Company           string `json:"company"`
	Phone             string `json:"phone"`
	Room_price        int    `json:"room_price"`
	Final_price       int    `json:"final_price"`
	User              User   `json:"user"`
	Room              Rooms  `json:"room"`
	Snack             Snacks `json:"snack"`
}

type ReservationInput struct {
	RoomID       string `json:"room_id" example:"6066f8a1-0a80-4299-86ca-99888912bbe5"`
	UserID       string `json:"user_id" example:"21691490-6817-4bf4-9bf7-3bf624d210a7"`
	SnackID      string `json:"snack_id" example:"b8f8cab4-9f0e-4d08-88aa-9fd465a52536"`
	StartTime    string `json:"start_time"`
	EndTime      string `json:"end_time"`
	BookingDate  string `json:"booking_date"`
	Participants int    `json:"participants"`
	Name         string `json:"name"`
	Total_Snack  int    `json:"total_snack"`
	Company      string `json:"company"`
	Phone        string `json:"phone"`
}

type ResponseReservation struct {
	ID                string `json:"id"`
	RoomID            string `json:"room_id"`
	UserID            string `json:"user_id"`
	SnackID           string `json:"snack_id"`
	StartTime         string `json:"start_time"`
	EndTime           string `json:"end_time"`
	BookingDate       string `json:"booking_date"`
	CreatedAt         string `json:"created_at"`
	UpdatedAt         string `json:"updated_at"`
	Status            string `json:"status"`
	Participants      int    `json:"participants"`
	Name              string `json:"name"`
	Total_Snack       int    `json:"total_snack"`
	Total_Snack_Price int    `json:"total_snack_price"`
	Company           string `json:"company"`
	Phone             string `json:"phone"`
	Room_price        int    `json:"room_price"`
	Final_price       int    `json:"final_price"`
}

var data = []Reservation{}

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

func getRoomPriceByID(roomID string) (int, error) {
	var roomPrice int
	query := "SELECT price FROM rooms WHERE id = $1"

	err := utils.DB.QueryRow(query, roomID).Scan(&roomPrice)
	if err != nil {
		return 0, err // Mengembalikan error jika room_id tidak ditemukan
	}

	return roomPrice, nil
}

func getSnackPriceByID(snackID string) (int, error) {
	var snackPrice int
	query := "SELECT price FROM snacks WHERE id = $1"

	err := utils.DB.QueryRow(query, snackID).Scan(&snackPrice)
	if err != nil {
		return 0, err // Mengembalikan error jika snack_id tidak ditemukan
	}

	return snackPrice, nil
}

func isRoomAvailable(roomID, bookingDate, startTime, endTime string) (bool, string) {
	loc, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		return false, "Failed to load location"
	}

	// Gabungkan bookingDate dan startTime menjadi satu string
	bookingDateTimeStr := fmt.Sprintf("%s %s", bookingDate, startTime)

	// Parse bookingDateTimeStr ke time.Time
	bookingDateTime, err := time.ParseInLocation("2006-01-02 15:04:05", bookingDateTimeStr, loc)
	if err != nil {
		return false, "Invalid date and time format"
	}

	// Cek apakah bookingDateTime adalah waktu lampau
	if bookingDateTime.Before(time.Now().In(loc)) {
		return false, "Waktu yang anda pilih sudah berlalu, Ayo move on melihat masa depan lebih cerah :D"
	}

	// Query untuk mengecek ketersediaan ruangan
	query := `
		SELECT COUNT(*) 
		FROM reservations 
		WHERE room_id = $1 
		AND booking_date = $2
		AND (
			(start_time <= $3 AND end_time > $3) OR
			(start_time < $4 AND end_time >= $4) OR
			(start_time >= $3 AND end_time <= $4)
		)
	`

	var count int
	err = utils.DB.QueryRow(query, roomID, bookingDate, startTime, endTime).Scan(&count)
	if err != nil {
		fmt.Println("Error checking room availability:", err)
		return false, "Error checking room availability"
	}

	// Jika count > 0, berarti ada konflik
	if count > 0 {
		return false, "Room is not available at the selected time"
	}

	return true, ""
}

// PostReservation godoc
// @Summary Create a new reservation
// @Description Note: Untuk Booking date menggunakan format YYYY-MM-DD
// @Tags reservations
// @Accept json
// @Produce json
// @Param reservation body ReservationInput true "Reservation details" Example({"room_id": "6066f8a1-0a80-4299-86ca-99888912bbe5", "user_id": "21691490-6817-4bf4-9bf7-3bf624d210a7", "snack_id": "b8f8cab4-9f0e-4d08-88aa-9fd465a52536", "start_time": "10:00:00", "end_time": "12:00:00", "booking_date": "2024-12-20", "name": "Sapto", "participants": 7, "total_snack": 7, "company": "PT. Sinau Koding Inc", "phone": "0191181811"})
// @Success 200 {object} map[string]string
// @Router /reservations [post]
func PostReservation(c echo.Context) error {
	var reservation ResponseReservation
	loc, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		fmt.Println("Error loading location:", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to load location"})
	}
	currentTime := time.Now().In(loc)
	formattedTime := currentTime.Format(time.RFC3339)

	// Bind input JSON ke struct Reservation
	if err := c.Bind(&reservation); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request payload"})
	}

	// Cek ketersediaan ruangan
	isAvailable, message := isRoomAvailable(reservation.RoomID, reservation.BookingDate, reservation.StartTime, reservation.EndTime)
	if !isAvailable {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": message})
	}

	roomPrice, err := getRoomPriceByID(reservation.RoomID)
	snackPrice, err := getSnackPriceByID(reservation.SnackID)
	if reservation.Total_Snack == 0 {
		reservation.Total_Snack = 0
	}
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to retrieve room price"})
	}
	reservation.Room_price = roomPrice
	reservation.Total_Snack_Price = reservation.Total_Snack * snackPrice
	reservation.Final_price = reservation.Room_price + reservation.Total_Snack_Price
	reservation.Status = "booked"
	reservation.CreatedAt = formattedTime
	reservation.UpdatedAt = formattedTime

	// Query untuk menyimpan reservasi ke database
	query := `
		INSERT INTO reservations (room_id, user_id, snack_id, start_time, end_time, participant, name, total_snack, total_snack_price, company, phone, room_price, final_price, booking_date, created_at, updated_at) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16) 
		RETURNING id
	`

	var reservationID string
	err = utils.DB.QueryRow(query, reservation.RoomID, reservation.UserID, reservation.SnackID, reservation.StartTime, reservation.EndTime, reservation.Participants, reservation.Name, reservation.Total_Snack, reservation.Total_Snack_Price, reservation.Company, reservation.Phone, reservation.Room_price, reservation.Final_price, reservation.BookingDate, formattedTime, formattedTime).Scan(&reservationID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	// Set ID yang dihasilkan ke dalam struct
	reservation.ID = reservationID

	// Kembalikan respons dengan data reservasi
	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message":     "Reservation created successfully",
		"reservation": reservation,
	})
}
