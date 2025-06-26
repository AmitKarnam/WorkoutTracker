package middleware

import (
	"net/http"
	"os"

	"github.com/AmitKarnam/WorkoutTracker/internal/constants"
	"github.com/AmitKarnam/WorkoutTracker/internal/models"
	"github.com/AmitKarnam/WorkoutTracker/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

const ErrTokenInvalid = "invalid access token"

func JWTValidate(userService services.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		accessToken, err := c.Cookie(constants.AccessTokenCookie)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "request should contain access token"})
			c.Abort()
			return
		}

		token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
			return os.Getenv("JWT_SIGNING"), nil
		}, jwt.WithValidMethods([]string{"HS256"}))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": ErrTokenInvalid})
			c.Abort()
			return
		}

		if !token.Valid {
			c.JSON(http.StatusBadRequest, gin.H{"error": ErrTokenInvalid})
			c.Abort()
			return
		}

		// extract and set jwt claims into context
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{"error": ErrTokenInvalid})
			c.Abort()
			return
		}

		// check the issuer
		if claims["iss"] != constants.Project {
			c.JSON(http.StatusBadRequest, gin.H{"error": ErrTokenInvalid})
			c.Abort()
			return
		}

		userID, ok := claims[constants.JWTUserID].(uint)
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{"error": ErrTokenInvalid})
			c.Abort()
			return
		}

		userRole, ok := claims[constants.JWTUserRole].(models.UserRole)
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{"error": ErrTokenInvalid})
			c.Abort()
			return
		}

		// verify the role passed from jwt to the one present in the database
		roleVerify, err := userService.VerifyUserRole(c.Request.Context(), userID, userRole)
		if err != nil || !roleVerify {
			c.JSON(http.StatusBadRequest, gin.H{"error": ErrTokenInvalid})
			c.Abort()
			return
		}

		c.Set(constants.JWTUserID, userID)
		c.Set(constants.JWTUserRole, userRole)

		c.Next()
	}
}
