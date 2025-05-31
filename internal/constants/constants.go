package constants

import (
	"time"
)

const (
	AccessTokenExpiry  = 15 * time.Minute
	RefreshTokenExpiry = 7 * 24 * time.Hour
	AccessTokenCookie  = "access-token"
	RefreshTokenCookie = "refresh-token"
	TokenPath          = "/"

	// Cookie security flags
	CookieSecure   = false // Set to true in production when using HTTPS
	CookieHttpOnly = true  // Prevents JavaScript access to cookies

	// JWT Claims
	JWTUserID   = "user_id"
	JWTUserRole = "user_role"
)
