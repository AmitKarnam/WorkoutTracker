package controllers

import (
	"net/http"

	"github.com/AmitKarnam/WorkoutTracker/models"

	"github.com/gin-gonic/gin"
)

type MuscleGroupController struct{}

func (mc *MuscleGroupController) Get(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": models.MuscleGroupList})
}

func (mc *MuscleGroupController) Post(c *gin.Context) {
	var newMuscleGroup models.MuscleGroup
	if err := c.ShouldBindBodyWithJSON(&newMuscleGroup); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	models.MuscleGroupList = append(models.MuscleGroupList, newMuscleGroup)
	c.JSON(http.StatusCreated, gin.H{"data": newMuscleGroup})
}
