package oauth

import "github.com/gin-gonic/gin"

type OAuthController interface {
	LoginHandler(c *gin.Context)
	CallbackHandler(c *gin.Context)
}
