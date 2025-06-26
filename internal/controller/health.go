package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthController struct{}

func (hc *HealthController) Get(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"Status": "Healthy"})
}
