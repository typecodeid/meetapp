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

type UserShow struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	ImageID  string `json:"image_id"`
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
	ID                string   `json:"id"`
	RoomID            string   `json:"room_id"`
	UserID            string   `json:"user_id"`
	SnackID           string   `json:"snack_id"`
	StartTime         string   `json:"start_time"`
	EndTime           string   `json:"end_time"`
	BookingDate       string   `json:"booking_date"`
	CreatedAt         string   `json:"created_at"`
	UpdatedAt         string   `json:"updated_at"`
	Status            string   `json:"status"`
	Participants      int      `json:"participants"`
	Name              string   `json:"name"`
	Total_Snack       int      `json:"total_snack"`
	Total_Snack_Price int      `json:"total_snack_price"`
	Company           string   `json:"company"`
	Phone             string   `json:"phone"`
	Room_price        int      `json:"room_price"`
	Final_price       int      `json:"final_price"`
	User              UserShow `json:"user"`
	Room              Rooms    `json:"room"`
	Snack             Snacks   `json:"snack"`
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

type UpdateStatusInput struct {
	Status string `json:"status" validate:"required" example:"cancel"`
}

var data = []Reservation{}

// GetAllReservation godoc
// @Summary Get all reservations
// @Description Retrieve all reservations in the system contoh: /reservations?status=cancel&room_type=medium
// @Tags reservations
// @Accept json
// @Produce json
// @Param status query string false "Filter by status"
// @Param room_type query string false "Filter by room type"
// @Param start_date query string false "Filter by start date (YYYY-MM-DD)"
// @Param end_date query string false "Filter by end date (YYYY-MM-DD)"
// @Success 200 {array} Reservation
// @Security BearerAuth
// @Router /reservations [get]
func GetAllReservation(c echo.Context) error {
	// Ambil query parameter
	status := c.QueryParam("status")
	roomType := c.QueryParam("room_type")
	startDateStr := c.QueryParam("start_date")
	endDateStr := c.QueryParam("end_date")

	// Inisialisasi query SQL dasar
	query := `
		SELECT 
			r.id, r.room_id, r.user_id, r.snack_id, r.start_time, r.end_time, r.booking_date, 
			r.created_at, r.updated_at, r.status, r.participant, r.name, r.total_snack, 
			r.total_snack_price, r.company, r.phone, r.room_price, r.final_price,
			rm.id AS room_id, rm.name AS room_name, rm.type AS room_type, rm.capacity AS room_capacity, rm.price AS room_price,
			u.id AS user_id, u.username AS user_username, u.email AS user_email, u.image_id AS user_image_id, u.role AS user_role, u.status AS user_status, u.language AS user_language,
			s.id AS snack_id, s.name AS snack_name, s.category AS snack_category, s.package AS snack_package, s.price AS snack_price
		FROM 
			reservations r
		LEFT JOIN 
			rooms rm ON r.room_id = rm.id
		LEFT JOIN 
			users u ON r.user_id = u.id
		LEFT JOIN 
			snacks s ON r.snack_id = s.id
		WHERE 
			1 = 1
	`

	// Variabel untuk menyimpan parameter
	params := []interface{}{}
	paramIndex := 1 // PostgreSQL menggunakan $1, $2, ...

	// Tambahkan filter berdasarkan query parameter jika tersedia
	if status != "" {
		query += fmt.Sprintf(" AND r.status = $%d", paramIndex)
		params = append(params, status)
		paramIndex++
	}

	if roomType != "" {
		query += fmt.Sprintf(" AND rm.type = $%d", paramIndex)
		params = append(params, roomType)
		paramIndex++
	}

	if startDateStr != "" {
		startDate, err := time.Parse("2006-01-02", startDateStr)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"message": "Invalid start date format",
			})
		}
		query += fmt.Sprintf(" AND r.booking_date >= $%d", paramIndex)
		params = append(params, startDate)
		paramIndex++
	}

	if endDateStr != "" {
		endDate, err := time.Parse("2006-01-02", endDateStr)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"message": "Invalid end date format",
			})
		}
		query += fmt.Sprintf(" AND r.booking_date <= $%d", paramIndex)
		params = append(params, endDate)
		paramIndex++
	}

	// Jalankan query dengan parameter
	rows, err := utils.DB.Query(query, params...)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}
	defer rows.Close()

	var reservations []Reservation

	// Iterasi hasil query dan mapping ke struct
	for rows.Next() {
		var reservationData Reservation

		err := rows.Scan(
			&reservationData.ID,
			&reservationData.RoomID,
			&reservationData.UserID,
			&reservationData.SnackID,
			&reservationData.StartTime,
			&reservationData.EndTime,
			&reservationData.BookingDate,
			&reservationData.CreatedAt,
			&reservationData.UpdatedAt,
			&reservationData.Status,
			&reservationData.Participants,
			&reservationData.Name,
			&reservationData.Total_Snack,
			&reservationData.Total_Snack_Price,
			&reservationData.Company,
			&reservationData.Phone,
			&reservationData.Room_price,
			&reservationData.Final_price,
			// Room details
			&reservationData.Room.ID,
			&reservationData.Room.Name,
			&reservationData.Room.Type,
			&reservationData.Room.Capacity,
			&reservationData.Room.Price,
			// User details
			&reservationData.User.ID,
			&reservationData.User.Username,
			&reservationData.User.Email,
			&reservationData.User.ImageID,
			&reservationData.User.Role,
			&reservationData.User.Status,
			&reservationData.User.Language,
			// Snack details
			&reservationData.Snack.ID,
			&reservationData.Snack.Name,
			&reservationData.Snack.Category,
			&reservationData.Snack.Package,
			&reservationData.Snack.Price,
		)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": err.Error(),
			})
		}
		reservations = append(reservations, reservationData)
	}

	dataResponse := APIResponse{
		Message: "Success",
		Status:  http.StatusOK,
		Data:    reservations,
	}

	return c.JSON(http.StatusOK, dataResponse)
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
// @Security BearerAuth
// @Router /reservations/{id} [get]
func GetByID(c echo.Context) error {
	id := c.Param("id")

	// Query untuk mendapatkan detail reservasi beserta user, room, dan snack tanpa kolom password
	query := `
		SELECT 
			r.id, r.room_id, r.user_id, r.snack_id, r.start_time, r.end_time, r.booking_date, 
			r.created_at, r.updated_at, r.status, r.participant, r.name, r.total_snack, 
			r.total_snack_price, r.company, r.phone, r.room_price, r.final_price,
			rm.id AS room_id, rm.name AS room_name, rm.type AS room_type, rm.capacity AS room_capacity, rm.price AS room_price,
			u.id AS user_id, u.username AS user_username, u.email AS user_email, u.image_id AS user_image_id, u.role AS user_role, u.status AS user_status, u.language AS user_language,
			s.id AS snack_id, s.name AS snack_name, s.category AS snack_category, s.package AS snack_package, s.price AS snack_price
		FROM 
			reservations r
		LEFT JOIN 
			rooms rm ON r.room_id = rm.id
		LEFT JOIN 
			users u ON r.user_id = u.id
		LEFT JOIN 
			snacks s ON r.snack_id = s.id
		WHERE 
			r.id = $1
	`

	var reservationData Reservation

	err := utils.DB.QueryRow(query, id).Scan(
		&reservationData.ID,
		&reservationData.RoomID,
		&reservationData.UserID,
		&reservationData.SnackID,
		&reservationData.StartTime,
		&reservationData.EndTime,
		&reservationData.BookingDate,
		&reservationData.CreatedAt,
		&reservationData.UpdatedAt,
		&reservationData.Status,
		&reservationData.Participants,
		&reservationData.Name,
		&reservationData.Total_Snack,
		&reservationData.Total_Snack_Price,
		&reservationData.Company,
		&reservationData.Phone,
		&reservationData.Room_price,
		&reservationData.Final_price,
		// Room details
		&reservationData.Room.ID,
		&reservationData.Room.Name,
		&reservationData.Room.Type,
		&reservationData.Room.Capacity,
		&reservationData.Room.Price,
		// User details (tanpa password)
		&reservationData.User.ID,
		&reservationData.User.Username,
		&reservationData.User.Email,
		&reservationData.User.ImageID,
		&reservationData.User.Role,
		&reservationData.User.Status,
		&reservationData.User.Language,
		// Snack details
		&reservationData.Snack.ID,
		&reservationData.Snack.Name,
		&reservationData.Snack.Category,
		&reservationData.Snack.Package,
		&reservationData.Snack.Price,
	)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "Reservation not found",
			"error":   err.Error(),
		})
	}

	dataResponse := APIResponse{
		Message: "Success",
		Status:  http.StatusOK,
		Data:    reservationData,
	}

	return c.JSON(http.StatusOK, dataResponse)
}

// PutReservation godoc
// @Summary Update reservation status
// @Description Update the status of a reservation by its ID
// @Tags reservations
// @Accept json
// @Produce json
// @Param id path string true "Reservation ID"
// @Param reservation body UpdateStatusInput true "Reservation status update"
// @Success 200 {object} map[string]interface{} "Success response"
// @Failure 400 {object} map[string]string "Invalid request payload or missing status"
// @Failure 404 {object} map[string]string "Reservation not found"
// @Failure 500 {object} map[string]string "Internal server error"
// @Security BearerAuth
// @Router /reservations/{id} [put]
func PutReservation(c echo.Context) error {
	// Ambil ID dari path parameter
	id := c.Param("id")

	var input UpdateStatusInput

	// Bind input JSON ke struct
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Invalid request payload",
		})
	}

	// Validasi input
	if input.Status == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Status is required",
		})
	}

	// Query untuk update status reservasi
	query := `
        UPDATE reservations
        SET status = $1, updated_at = NOW()
        WHERE id = $2
        RETURNING id, status
    `

	var updatedID, updatedStatus string

	err := utils.DB.QueryRow(query, input.Status, id).Scan(&updatedID, &updatedStatus)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Failed to update reservation",
			"error":   err.Error(),
		})
	}
	dataResponse := APIResponse{
		Message: "Reservation updated successfully",
		Status:  http.StatusOK,
		id:      updatedID,
	}

	return c.JSON(http.StatusOK, dataResponse)
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
		return false, "Kenangan kadang sulit dilupakan. Ayo move on, tanggal ini sudah lewat"
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
// @Description Note: Untuk Booking date menggunakan format YYYY-MM-DD untuk time menggunakan format HH:MM:SS
// @Tags reservations
// @Accept json
// @Produce json
// @Param reservation body ReservationInput true "Reservation details" Example({"room_id": "6066f8a1-0a80-4299-86ca-99888912bbe5", "user_id": "21691490-6817-4bf4-9bf7-3bf624d210a7", "snack_id": "b8f8cab4-9f0e-4d08-88aa-9fd465a52536", "start_time": "10:00:00", "end_time": "12:00:00", "booking_date": "2024-12-20", "name": "Sapto", "participants": 7, "total_snack": 7, "company": "PT. Sinau Koding Inc", "phone": "0191181811"})
// @Success 200 {object} map[string]string
// @Security BearerAuth
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

	dataResponse := APIResponse{
		Message: "Reservation created successfully",
		Status:  http.StatusCreated,
		Data:    reservation,
	}

	// Kembalikan respons dengan data reservasi
	return c.JSON(http.StatusCreated, dataResponse)
}
