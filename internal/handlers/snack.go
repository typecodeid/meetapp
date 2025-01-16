package handlers

import (
	utils "meetapp/pkg/database"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Snack struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Category string `json:"category"`
	Package  string `json:"package"`
	Price    string `json:"price"`
}

type responseSnack struct {
	Message string  `json:"message"`
	Data    []Snack `json:"data"`
}

// GetSnack godoc
// @Summary Get Snack
// @Description Get Snack
// @Tags Snack
// @Produce json
// @Success 200 {object} map[string]string
// @Security BearerAuth
// @Router /snack [get]
func GetSnack(c echo.Context) error {
	query := "SELECT id, name, category, package, price FROM snacks"

	rows, err := utils.DB.Query(query)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to retrieve snacks",
		})
	}
	var snack []Snack

	for rows.Next() {
		var snackData Snack
		err := rows.Scan(&snackData.ID, &snackData.Name, &snackData.Category, &snackData.Package, &snackData.Price)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": "Failed to scan snack data",
			})
		}
		snack = append(snack, snackData)
	}

	response := responseSnack{
		Message: "Success",
		Data:    snack,
	}
	return c.JSON(http.StatusOK, response)
}

func GetSnackByID(c echo.Context) error {
	id := c.Param("id")
	query := "SELECT id, name, type, category, package, price FROM snacks WHERE id = $1"
	var snack Snack

	err := utils.DB.QueryRow(query, id).Scan(&snack.ID, &snack.Name, &snack.Category, &snack.Price)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to retrieve snack",
		})
	}
	response := responseSnack{
		Message: "Success",
		Data:    []Snack{snack},
	}
	return c.JSON(http.StatusOK, response)
}

// CreateSnack godoc
// @Summary Create Snack
// @Description Create Snack
// @Tags Snack
// @Produce json
// @Success 200 {object} map[string]string
// @Security BearerAuth
// @Router /snack [post]
func CreateSnack(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Create Snack Here",
	})
}
