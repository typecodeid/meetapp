package mymiddleware

import (
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

var jwtSecret = []byte("ScreetKey007") // Ganti dengan secret key Anda

// Middleware untuk validasi token dan role-based access control
func TokenRole(requiredRole string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Ambil token dari header Authorization
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return c.JSON(http.StatusUnauthorized, map[string]string{
					"message": "Missing or invalid token",
				})
			}

			// Pastikan token dimulai dengan "Bearer "
			if !strings.HasPrefix(authHeader, "Bearer ") {
				return c.JSON(http.StatusUnauthorized, map[string]string{
					"message": "Invalid token format",
				})
			}

			// Ambil token (menghilangkan "Bearer ")
			tokenString := strings.TrimPrefix(authHeader, "Bearer ")

			// Parse dan validasi token
			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, echo.NewHTTPError(http.StatusUnauthorized, "Unexpected signing method")
				}
				return jwtSecret, nil
			})

			if err != nil {
				return c.JSON(http.StatusUnauthorized, map[string]string{
					"message": "Invalid or expired token",
				})
			}

			// Ambil klaim dari token
			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok || !token.Valid {
				return c.JSON(http.StatusUnauthorized, map[string]string{
					"message": "Invalid token claims",
				})
			}

			// Ambil role dari klaim
			role, ok := claims["role"].(string)
			if !ok {
				return c.JSON(http.StatusUnauthorized, map[string]string{
					"message": "Invalid token payload: role missing",
				})
			}

			// Periksa apakah role sesuai dengan requiredRole
			if requiredRole != "" && !(role == requiredRole || (requiredRole == "user" && role == "admin")) {
				return c.JSON(http.StatusForbidden, map[string]string{
					"message": "Forbidden: insufficient permissions",
				})
			}

			// Set klaim ke context untuk digunakan di handler berikutnya
			c.Set("userRole", role)
			c.Set("userEmail", claims["email"])

			// Lanjutkan ke handler berikutnya
			return next(c)
		}
	}
}
