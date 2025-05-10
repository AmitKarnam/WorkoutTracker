package jwt

import (
	"fmt"
	"os"
	"time"

	"github.com/AmitKarnam/WorkoutTracker/internal/constants"
	"github.com/AmitKarnam/WorkoutTracker/internal/models"
	"github.com/golang-jwt/jwt/v4"
)

func GenerateUserAccessToken(user *models.User) (string, error) {
	jwtSecret := os.Getenv("JWT_SIGNING")

	claims := jwt.MapClaims{
		"sub":     fmt.Sprintf("%d", user.ID),
		"user_id": user.ID,
		"role":    user.Role,
		"iat":     time.Now().Unix(),
		"exp":     time.Now().Add(constants.AccessTokenExpiry).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(jwtSecret))
}
