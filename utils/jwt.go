package utils

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var jwtKey = []byte("secret") // Ganti dengan value dari os.Getenv("JWT_SECRET") di production

// Generate token JWT
func GenerateJWT(adminID int64) (string, error) {
	claims := jwt.MapClaims{
		"admin_id": adminID,
		"exp":      time.Now().Add(24 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

// Parse token JWT dan ambil admin_id
func ParseJWT(tokenString string) (int64, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil || !token.Valid {
		return 0, errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("invalid claims")
	}

	adminIDFloat, ok := claims["admin_id"].(float64)
	if !ok {
		return 0, errors.New("admin_id not found in token")
	}

	return int64(adminIDFloat), nil
}

// Ekstrak admin_id dari header Authorization: Bearer <token>
func ExtractAdminIDFromRequest(r *http.Request) (int64, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return 0, errors.New("authorization header not found")
	}

	tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
	return ParseJWT(tokenStr)
}
