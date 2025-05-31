package controller

import (
	"net/http"

	"github.com/AmitKarnam/WorkoutTracker/internal/constants"
	"github.com/AmitKarnam/WorkoutTracker/internal/services"
	"github.com/AmitKarnam/WorkoutTracker/logger"
	"github.com/gin-gonic/gin"
)

type LogoutController interface {
	Logout(c *gin.Context)
}

type logoutController struct {
	service services.RefreshTokenService
}

func NewLogoutController() LogoutController {
	return &logoutController{}
}

func (lc *logoutController) Logout(c *gin.Context) {
	ctx := c.Request.Context()

	logger.Logger.Info("recieved request to logout user")

	refreshToken, err := c.Cookie(constants.RefreshTokenCookie)
	if err != nil {
		logger.Logger.Error("error retriveing refresh token from request", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "refresh token not found in request"})
		return
	}

	err = lc.service.Delete(ctx, refreshToken)
	if err != nil {
		logger.Logger.Error("error deleting refresh token from database", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to delete refresh token"})
		return
	}

	// Clear cookies
	c.SetCookie(constants.AccessTokenCookie, "", -1, constants.TokenPath, "", constants.CookieSecure, constants.CookieHttpOnly)
	c.SetCookie(constants.RefreshTokenCookie, "", -1, constants.TokenPath, "", constants.CookieSecure, constants.CookieHttpOnly)

	c.JSON(http.StatusOK, gin.H{"message": "logged out successfully"})
}
