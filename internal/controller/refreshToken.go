package controller

import (
	"net/http"
	"os"

	"github.com/AmitKarnam/WorkoutTracker/internal/constants"
	"github.com/AmitKarnam/WorkoutTracker/internal/services"
	"github.com/AmitKarnam/WorkoutTracker/logger"
	"github.com/gin-gonic/gin"
)

type RefreshTokenController interface {
	RefreshTokenHandler(c *gin.Context)
}

type refreshTokenController struct {
	service services.RefreshTokenService
}

func NewRefreshTokenController(service services.RefreshTokenService) RefreshTokenController {
	return &refreshTokenController{service: service}
}

func (rtc *refreshTokenController) RefreshTokenHandler(c *gin.Context) {
	ctx := c.Request.Context()
	logger.Logger.Info("recieved request to refresh the tokens")

	// get refresh token from cookies in request
	refreshToken, err := c.Cookie(constants.RefreshTokenCookie)
	if err != nil {
		logger.Logger.Error("refresh token not found in request", "error", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "refresh token not found"})
		return
	}

	newaccessToken, newrefreshToken, err := rtc.service.HandleRefresh(ctx, refreshToken)
	if err != nil {
		logger.Logger.Error("unable to refresh tokens", "error", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unable to refresh token"})
		return
	}

	c.SetCookie(constants.AccessTokenCookie, newaccessToken, int(constants.AccessTokenExpiry.Seconds()), constants.TokenPath, os.Getenv("DOMAIN"), constants.CookieSecure, constants.CookieHttpOnly)
	c.SetCookie(constants.RefreshTokenCookie, newrefreshToken, int(constants.RefreshTokenExpiry.Seconds()), constants.TokenPath, os.Getenv("DOMAIN"), constants.CookieSecure, constants.CookieHttpOnly)

	logger.Logger.Info("refresh token successfully refreshed")
	c.JSON(http.StatusOK, gin.H{"message": "tokens successfully refreshed"})
}
