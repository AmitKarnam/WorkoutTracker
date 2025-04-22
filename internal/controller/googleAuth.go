package controller

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"net/http"

	"github.com/AmitKarnam/WorkoutTracker/internal/services"
	"github.com/gin-gonic/gin"
)

var stateToken string // Temporary storage for the state token (TODO : Need to use session based token management)

type AuthController struct {
	authService services.AuthService
}

func NewAuthController(authService services.AuthService) *AuthController {
	return &AuthController{authService: authService}
}

// Generate a random state token
func generateStateToken() string {
	b := make([]byte, 16)
	rand.Read(b)
	return base64.URLEncoding.EncodeToString(b)
}

func (ac *AuthController) GoogleLogin(c *gin.Context) {
	// Generate and store the state token
	stateToken = generateStateToken()

	// Pass the state token to the OAuth URL
	url := ac.authService.GetGoogleOAuthURL(stateToken)
	c.Redirect(http.StatusFound, url)
}

func (ac *AuthController) GoogleCallback(c *gin.Context) {
	// Retrieve the state and code from the query parameters
	state := c.Query("state")
	code := c.Query("code")

	// Validate the state token
	if state != stateToken {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid state token"})
		return
	}

	if code == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "code not provided"})
		return
	}

	// Handle the callback
	user, err := ac.authService.HandleGoogleCallback(context.Background(), code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Generate JWT for successful user creation

	c.JSON(http.StatusOK, gin.H{"user": user})
}
