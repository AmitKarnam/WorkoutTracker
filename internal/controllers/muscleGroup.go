package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/AmitKarnam/WorkoutTracker/internal/models"
	"github.com/AmitKarnam/WorkoutTracker/internal/services"
	"github.com/AmitKarnam/WorkoutTracker/logger"
	"github.com/gin-gonic/gin"
)

type MuscleGroupController interface {
	Get(*gin.Context)
	GetByID(*gin.Context)
	Post(*gin.Context)
	Put(*gin.Context)
	Delete(*gin.Context)
}

type muscleGroupController struct {
	service services.MuscleGroupService
}

func NewMuscleGroupController(service services.MuscleGroupService) MuscleGroupController {
	return &muscleGroupController{service: service}
}

// Get method to fetch all muscle groups from database
func (msc *muscleGroupController) Get(c *gin.Context) {
	logger.Logger.Info("received request to get all muscle groups")
	muscleGroups, err := msc.service.GetAll()
	if err != nil {
		logger.Logger.Error("error fetching muscle groups", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error fetching muscle groups"})
		return
	}
	logger.Logger.Info("successfully fetched all muscle groups")
	c.JSON(http.StatusOK, gin.H{"data": muscleGroups})
}

// GetByID method to fetch a muscle group by ID from database
func (msc *muscleGroupController) GetByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		logger.Logger.Error("error fetching muscle group by id, string converstion to interger error", "id", id)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid muscle group ID"})
		return
	}
	logger.Logger.Info("received request to fetch muscle group by id", "id", id)

	muscleGroup, err := msc.service.GetByID(uint(id))
	if err != nil {
		logger.Logger.Error("error fetching muscle group by id", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("error fetching muscle group by id, %s", err.Error())})
		return
	}
	logger.Logger.Info("successfully fetched muscle group by id", "id", id)
	c.JSON(http.StatusOK, gin.H{"data": muscleGroup})
}

// Post method to add a new muscle group to database
func (msc *muscleGroupController) Post(c *gin.Context) {
	logger.Logger.Info("received request to create new muscle group")
	var muscleGroup models.MuscleGroup
	if err := c.ShouldBindJSON(&muscleGroup); err != nil {
		logger.Logger.Error("error creating muscle group, error parsing request body", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	if strings.TrimSpace(muscleGroup.MuscleGroup) == "" {
		logger.Logger.Error("error creating muscle group, muscle group cannot be empty")
		c.JSON(http.StatusBadRequest, gin.H{"error": "muscle group name is required"})
		return
	}

	muscleGroup.MuscleGroup = strings.ToLower(muscleGroup.MuscleGroup)

	if err := msc.service.Create(&muscleGroup); err != nil {
		logger.Logger.Error("error creating muscle group, error saving muscle group to database", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("error creating muscle group, %s", err.Error())})
		return
	}
	logger.Logger.Info("successfully created muscle group", "name", muscleGroup.MuscleGroup)
	c.JSON(http.StatusCreated, gin.H{"data": muscleGroup})
}

// Put method to edit a muscle group record
func (msc *muscleGroupController) Put(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		logger.Logger.Error("error updating muscle group, string converstion to interger error", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid muscle group ID"})
		return
	}
	logger.Logger.Info("received request to update muscle group by id", "id", id)

	var input models.MuscleGroup
	if err := c.ShouldBindJSON(&input); err != nil {
		logger.Logger.Error("error updating muscle group, error parsing request body", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	if strings.TrimSpace(input.MuscleGroup) == "" {
		logger.Logger.Error("error updating muscle group, muscle group cannot be empty")
		c.JSON(http.StatusBadRequest, gin.H{"error": "muscle group name is required"})
		return
	}

	input.MuscleGroup = strings.ToLower(input.MuscleGroup)

	muscleGroup, err := msc.service.Update(uint(id), input)
	if err != nil {
		logger.Logger.Error("error updating muscle group in database", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("error updating muscle group, %s", err.Error())})
		return
	}
	logger.Logger.Info("successfully updated muscle group by id", "id", id)
	c.JSON(http.StatusOK, gin.H{"data": muscleGroup})
}

// Delete method to delete an existing muscle group from database
func (msc *muscleGroupController) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		logger.Logger.Error("error deleting muscle group, string converstion to interger error", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid muscle group ID"})
		return
	}
	logger.Logger.Info("received request to delete muscle group by id", "id", id)

	if err := msc.service.Delete(uint(id)); err != nil {
		logger.Logger.Error("error deleting muscle group", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("error deleting muscle group, %s", err.Error())})
		return
	}
	logger.Logger.Info("successfully deleted muscle group by id", "id", id)
	c.JSON(http.StatusOK, gin.H{"message": "muscle group deleted successfully"})
}
