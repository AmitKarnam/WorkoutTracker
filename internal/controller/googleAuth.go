package controller

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"net/http"
	"os"
	"sync"

	"github.com/AmitKarnam/WorkoutTracker/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var sessionState = make(map[string]string)
var mutex = &sync.Mutex{}

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
	stateToken := generateStateToken()
	session_id := uuid.New().String()

	// Protect the sessionState map with a mutex
	mutex.Lock()
	sessionState[session_id] = stateToken
	mutex.Unlock()

	// TODO : the values for setting a cookie should be fetched from constant
	c.SetCookie("google_auth_session", session_id, 3600, "/", os.Getenv("DOMAIN"), false, true)

	// Pass the state token to the OAuth URL
	url := ac.authService.GetGoogleOAuthURL(stateToken)
	c.Redirect(http.StatusFound, url)
}

func (ac *AuthController) GoogleCallback(c *gin.Context) {
	// Retrieve the state and code from the query parameters
	state := c.Query("state")
	code := c.Query("code")

	// Get session_id from the request object, that is passed as http-only cookie
	session_id, err := c.Cookie("google_auth_session")
	fmt.Println("session id: ", session_id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Cookie not found"})
		return
	}

	// Protect the sessionState map with a mutex
	mutex.Lock()
	session_state, exists := sessionState[session_id]
	if exists {
		delete(sessionState, session_id)
	}
	mutex.Unlock()

	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "session id does not exsist"})
		return
	}

	// Validate the state token
	if state != session_state {
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
