package controllers

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/AmitKarnam/WorkoutTracker/internal/models"
	"github.com/AmitKarnam/WorkoutTracker/internal/services"
	"github.com/AmitKarnam/WorkoutTracker/logger"
	"github.com/gin-gonic/gin"
)

type MuscleGroupController struct {
	service services.MuscleGroupService
}

func NewMuscleGroupController(service services.MuscleGroupService) *MuscleGroupController {
	return &MuscleGroupController{service: service}
}

// Get method to fetch all muscle groups from database
func (msc *MuscleGroupController) Get(c *gin.Context) {
	muscleGroups, err := msc.service.GetAll()
	if err != nil {
		logger.Logger.Error("error fetching muscle groups", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error fetching muscle groups"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": muscleGroups})
}

// GetByID method to fetch a muscle group by ID from database
func (msc *MuscleGroupController) GetByID(c *gin.Context) {
	idStr := c.Query("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid muscle group ID"})
		return
	}

	muscleGroup, err := msc.service.GetByID(uint(id))
	if err != nil {
		logger.Logger.Error("error fetching muscle group", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error fetching muscle group"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": muscleGroup})
}

// Post method to add a new muscle group to database
func (msc *MuscleGroupController) Post(c *gin.Context) {
	var muscleGroup models.MuscleGroup
	if err := c.ShouldBindJSON(&muscleGroup); err != nil {
		logger.Logger.Error("error parsing request body", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	muscleGroup.MuscleGroup = strings.ToLower(muscleGroup.MuscleGroup)

	if err := msc.service.Create(&muscleGroup); err != nil {
		logger.Logger.Error("error saving muscle group to database", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error saving muscle group"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": muscleGroup})
}

// Put method to edit a muscle group record
func (msc *MuscleGroupController) Put(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		logger.Logger.Error("error converting muscle group id to integer", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid muscle group ID"})
		return
	}

	var input models.MuscleGroup
	if err := c.ShouldBindJSON(&input); err != nil {
		logger.Logger.Error("error parsing request body", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	input.MuscleGroup = strings.ToLower(input.MuscleGroup)

	muscleGroup, err := msc.service.Update(uint(id), input)
	if err != nil {
		logger.Logger.Error("error updating muscle group in database", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error updating muscle group"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": muscleGroup})
}

// Delete method to delete an existing muscle group from database
func (msc *MuscleGroupController) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		logger.Logger.Error("error converting muscle group id to integer", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid muscle group ID"})
		return
	}

	if err := msc.service.Delete(uint(id)); err != nil {
		logger.Logger.Error("error deleting muscle group", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error deleting muscle group"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "muscle group deleted successfully"})
}
