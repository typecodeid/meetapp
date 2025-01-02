package handlers

import (
	utils "meetapp/pkg/database"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Dahsboard godoc
// @Summary Get dashboard data
// @Description Get dashboard data
// @Tags dashboard
// @Produce json
// @Router /dashboard [get]
func GetDashboard(c echo.Context) error {
	type RoomStats struct {
		RoomName        string  `json:"room_name"`
		UsagePercentage float64 `json:"usage_percentage"`
		RoomRevenue     int     `json:"room_revenue"`
	}

	type DashboardResponse struct {
		TotalRevenue      int         `json:"total_revenue"`
		TotalReservations int         `json:"total_reservations"`
		TotalVisitors     int         `json:"total_visitors"`
		TotalRooms        int         `json:"total_rooms"`
		RoomStats         []RoomStats `json:"room_stats"`
	}

	// Query for total revenue
	var totalRevenue int
	err := utils.DB.QueryRow(`
		SELECT COALESCE(SUM(final_price), 0)
		FROM reservations
		WHERE status = 'paid'
	`).Scan(&totalRevenue)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to fetch total revenue"})
	}

	// Query for total reservations
	var totalReservations int
	err = utils.DB.QueryRow(`
		SELECT COALESCE(SUM(participant), 0)
		FROM reservations
	`).Scan(&totalReservations)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to fetch total reservations"})
	}

	// Query for total visitors
	var totalVisitors int
	err = utils.DB.QueryRow(`
		SELECT COUNT(DISTINCT user_id)
		FROM reservations
	`).Scan(&totalVisitors)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to fetch total visitors"})
	}

	// Query for total rooms
	var totalRooms int
	err = utils.DB.QueryRow(`
		SELECT COUNT(*)
		FROM rooms
	`).Scan(&totalRooms)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to fetch total rooms"})
	}

	// Query for room stats
	rows, err := utils.DB.Query(`
		SELECT 
			r.name AS room_name,
			COALESCE(ROUND(COUNT(res.id)::decimal / NULLIF((SELECT COUNT(*) FROM reservations), 0) * 100, 2), 0) AS usage_percentage,
			COALESCE(SUM(res.final_price), 0) AS room_revenue
		FROM rooms r
		LEFT JOIN reservations res ON r.id = res.room_id AND res.status = 'paid'
		GROUP BY r.id, r.name
		ORDER BY r.name
	`)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to fetch room stats"})
	}
	defer rows.Close()

	var roomStats []RoomStats
	for rows.Next() {
		var stat RoomStats
		if err := rows.Scan(&stat.RoomName, &stat.UsagePercentage, &stat.RoomRevenue); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to parse room stats"})
		}
		roomStats = append(roomStats, stat)
	}

	// Construct response
	response := DashboardResponse{
		TotalRevenue:      totalRevenue,
		TotalReservations: totalReservations,
		TotalVisitors:     totalVisitors,
		TotalRooms:        totalRooms,
		RoomStats:         roomStats,
	}

	return c.JSON(http.StatusOK, response)
}
