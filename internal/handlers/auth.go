package handlers

import (
	utils "meetapp/pkg/database"
	"net/http"
	"strings"

	// "github.com/go-playground/validator/v10"
	// "github.com/google/uuid"
	"github.com/go-playground/validator"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	// "golang.org/x/crypto/bcrypt"
)

type UserInput struct {
	Username        string `json:"username" validate:"required"`
	Email           string `json:"email" validate:"required,email"`
	Password        string `json:"password" validate:"required,min=6"`
	ConfirmPassword string `json:"confirm_password" validate:"required,eqfield=Password"`
}

type UserResponse struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

var validate = validator.New()

func AuthLogin(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success login",
	})
}

func AuthRegister(c echo.Context) error {
	var input UserInput

	// Bind input JSON ke UserInput struct
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request payload"})
	}

	// Validasi input
	if err := validate.Struct(input); err != nil {
		// Kumpulkan error validasi
		var errors []string
		for _, err := range err.(validator.ValidationErrors) {
			errors = append(errors, strings.Title(err.Field())+" is "+err.Tag())
		}
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Validation failed",
			"errors":  errors,
		})
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to hash password"})
	}

	// Generate unique ID untuk pengguna
	userID := uuid.New().String()

	// Atur nilai default untuk field yang tidak diinput
	defaultImageID := "/images/no-image"
	defaultRole := "user"
	defaultStatus := false
	defaultLanguage := "id"

	// Persiapkan query SQL dengan nilai default
	query := `INSERT INTO users (id, image_id, username, email, password, role, status, language)
	          VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`

	// Eksekusi query dengan parameter
	_, err = utils.DB.Exec(query, userID, defaultImageID, input.Username, input.Email, string(hashedPassword), defaultRole, defaultStatus, defaultLanguage)
	if err != nil {
		// Periksa pelanggaran unique constraint (misalnya, email atau username sudah ada)
		if strings.Contains(err.Error(), "unique") {
			return c.JSON(http.StatusBadRequest, map[string]string{"message": "Username or Email already exists"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to register user"})
	}

	// Siapkan response tanpa informasi sensitif
	userResponse := UserResponse{
		ID:       userID,
		Username: input.Username,
		Email:    input.Email,
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Successfully registered",
		"user":    userResponse,
	})
}
