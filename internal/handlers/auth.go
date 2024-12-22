package handlers

import (
	"database/sql"
	utils "meetapp/pkg/database"
	"net/http"
	"strings"
	"time"

	"github.com/go-playground/validator"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
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
	Password string `json:"password"`
	Role     string `json:"role"`
	Status   bool   `json:"status"`
	Language string `json:"language"`
	ImageID  string `json:"image_id"`
}

type UserLogin struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

var jwtSecret = []byte("ScreetKey007")

var validate = validator.New()

// AuthLogin godoc
// @Summary Login user
// @Description Login user menggunakan email dan password, email: mail@mail.com, password: password123
// @Tags auth
// @Accept json
// @Produce json
// @Param user body UserLogin true "User login details"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string "Invalid request payload"
// @Failure 401 {object} map[string]string "Invalid credentials"
// @Router /login [post]
func AuthLogin(c echo.Context) error {
	var input UserLogin

	// Bind input JSON ke struct UserLogin
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request payload"})
	}

	// Query untuk mencari user berdasarkan email
	query := "SELECT id, email, password, username, role FROM users WHERE email = $1 AND status = true"
	var user User
	err := utils.DB.QueryRow(query, input.Email).Scan(&user.ID, &user.Email, &user.Password, &user.Username, &user.Role)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Invalid credentials"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to retrieve user"})
	}

	// Verifikasi password dengan bcrypt
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Invalid credentials"})
	}

	// Generate JWT token berdasarkan role
	token, err := generateJWT(user.Email, user.Role)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to generate token"})
	}

	// Kirim response dengan token
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Login successful",
		"token":   token,
		"role":    user.Role,
	})
}

// generateJWT membuat token JWT untuk pengguna
func generateJWT(email, role string) (string, error) {
	// Buat klaim token
	claims := jwt.MapClaims{
		"email": email,
		"role":  role,
		"exp":   time.Now().Add(time.Hour * 24).Unix(), // Token valid selama 24 jam
	}

	// Buat token menggunakan klaim
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Tanda tangani token dengan secret key
	signedToken, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

// AuthRegister godoc
// @Summary Register a new user
// @Description Register a new user
// @Tags auth
// @Accept json
// @Produce json
// @Param user body UserInput true "User details"
// @Success 200 {object} map[string]interface{}
// @Router /register [post]
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
