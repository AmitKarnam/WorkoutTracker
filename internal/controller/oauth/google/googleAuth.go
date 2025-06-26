package google

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"net/http"
	"os"
	"sync"

	"github.com/AmitKarnam/WorkoutTracker/internal/auth/jwt"
	"github.com/AmitKarnam/WorkoutTracker/internal/constants"
	"github.com/AmitKarnam/WorkoutTracker/internal/controller/oauth"
	"github.com/AmitKarnam/WorkoutTracker/internal/services"
	"github.com/AmitKarnam/WorkoutTracker/logger"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var sessionState = make(map[string]string)
var mutex = &sync.Mutex{}

type GoogleAuthController struct {
	googleAuthService   services.GoogleAuthService
	refreshTokenService services.RefreshTokenService
}

func NewGoogleAuthController(authService services.GoogleAuthService, refreshTokenService services.RefreshTokenService) oauth.OAuthController {
	return &GoogleAuthController{googleAuthService: authService, refreshTokenService: refreshTokenService}
}

// Generate a random state token
func generateStateToken() string {
	b := make([]byte, 16)
	rand.Read(b)
	return base64.URLEncoding.EncodeToString(b)
}

func (ac *GoogleAuthController) LoginHandler(c *gin.Context) {

	logger.Logger.Info("received google login request")
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
	url := ac.googleAuthService.GetGoogleOAuthURL(stateToken)
	logger.Logger.Info("redirecting request to google server")
	c.Redirect(http.StatusFound, url)
}

func (ac *GoogleAuthController) CallbackHandler(c *gin.Context) {
	ctx := c.Request.Context()
	// Retrieve the state and code from the query parameters
	state := c.Query("state")
	code := c.Query("code")

	// Get session_id from the request object, that is passed as http-only cookie
	session_id, err := c.Cookie("google_auth_session")
	if err != nil {
		logger.Logger.Info("error getting session-id from request", "error", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "cookie not found"})
		return
	}

	// Protect the sessionState map with a mutex
	mutex.Lock()
	session_state, exists := sessionState[session_id]
	if exists {
		delete(sessionState, session_id)
	}
	mutex.Unlock()
	logger.Logger.Info("deleted the session-id from the session_state")

	if !exists {
		logger.Logger.Error("session-id not found in session_state")
		c.JSON(http.StatusNotFound, gin.H{"error": "session id does not exsist"})
		return
	}

	// Validate the state token
	if state != session_state {
		logger.Logger.Error("seesion-id from the request is not valid")
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid state token"})
		return
	}

	if code == "" {
		logger.Logger.Error("auth code not provided in request")
		c.JSON(http.StatusBadRequest, gin.H{"error": "auth code not provided"})
		return
	}

	// Handle the callback
	user, err := ac.googleAuthService.HandleGoogleCallback(context.Background(), code)
	if err != nil {
		logger.Logger.Error("error handling the callback request", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	accessToken, err := jwt.GenerateUserAccessToken(user)
	if err != nil {
		logger.Logger.Error("error creating jwt user access token", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating tokens"})
		return
	}

	refreshToken, err := ac.refreshTokenService.CreateOrUpdate(ctx, user.ID)
	if err != nil {
		logger.Logger.Error("error creating jwt user refresh token", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating tokens"})
		return
	}

	c.SetCookie(constants.AccessTokenCookie, accessToken, int(constants.AccessTokenExpiry.Seconds()), constants.TokenPath, os.Getenv("DOMAIN"), constants.CookieSecure, constants.CookieHttpOnly)
	c.SetCookie(constants.RefreshTokenCookie, refreshToken, int(constants.RefreshTokenExpiry.Seconds()), constants.TokenPath, os.Getenv("DOMAIN"), constants.CookieSecure, constants.CookieHttpOnly)

	logger.Logger.Info("successfully authenticated the user")
	c.JSON(http.StatusOK, gin.H{"user": user})
}
