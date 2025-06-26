package middleware

import (
	"github.com/AmitKarnam/WorkoutTracker/internal/constants"
	"github.com/AmitKarnam/WorkoutTracker/internal/models"
	"github.com/gin-gonic/gin"
)

const ErrOperationNotPermitted = "operation not permitted for the user role"

// RBAC is a middleware function that checks if the user has the required role to access a route.
// If the user's role does not match the allowed role, it returns an error response.
// If the user's role matches, it allows the request to proceed.
// It expects the user's role to be set in the context with the key constants.JWTUserRole

func RBAC(allowedRole models.UserRole) gin.HandlerFunc {
	return func(c *gin.Context) {
		if allowedRole != c.Value(constants.JWTUserRole) {
			c.JSON(403, gin.H{"error": ErrOperationNotPermitted})
			c.Abort()
			return
		}

		c.Next()
	}
}
