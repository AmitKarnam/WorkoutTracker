package controller

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

type StrengthExerciseController interface {
	Get(c *gin.Context)
	GetByID(c *gin.Context)
	Post(c *gin.Context)
	Put(c *gin.Context)
	Delete(c *gin.Context)
}

type strengthExerciseController struct {
	service services.StrengthExerciseService
}

func NewStrengthExerciseController(service services.StrengthExerciseService) StrengthExerciseController {
	return &strengthExerciseController{service: service}
}

// Get All Strength Exercises from database
func (sec *strengthExerciseController) Get(c *gin.Context) {
	logger.Logger.Info("recieved request to Get all strength exercise")
	ctx := c.Request.Context()
	muscleGroup := c.Query("muscle-group")

	if muscleGroup != "" {
		logger.Logger.Info("recieved specific strength exercise query")
		strengthExercises, err := sec.service.GetByMuscleGroup(ctx, muscleGroup)
		if err != nil {
			logger.Logger.Error(fmt.Sprintf("error fetching strength exercises based on particular strength exercise %s", muscleGroup), "error", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error fetching strength exercise"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": strengthExercises})
		return
	}

	strengthExercises, err := sec.service.GetAll(ctx)
	if err != nil {
		logger.Logger.Error("error fetching strength exercises", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error fetching strength exercises"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": strengthExercises})
}

// Get sepcific Strength Exercise by ID from database
func (sec *strengthExerciseController) GetByID(c *gin.Context) {
	ctx := c.Request.Context()
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		logger.Logger.Error("error updating strength exercises, string converstion to interger error", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid strength exercise ID"})
		return
	}
	logger.Logger.Info("received request to update strength exercise by id", "id", id)

	strengthExercise, err := sec.service.GetByID(ctx, uint(id))
	if err != nil {
		logger.Logger.Error("error fetching strength exercise by id", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("error fetching strength exercise by id, %s", err.Error())})
		return
	}
	logger.Logger.Info("successfully fetched strength exercise by id", "id", id)
	c.JSON(http.StatusOK, gin.H{"data": strengthExercise})
}

// Post method to add a new Strength Exercise to database
func (sec *strengthExerciseController) Post(c *gin.Context) {
	logger.Logger.Info("received request to create strength exercise")
	ctx := c.Request.Context()
	var strengthExercise models.StrengthExercise
	err := c.ShouldBindBodyWithJSON(&strengthExercise)
	if err != nil {
		logger.Logger.Error("error creating strength exercise, error parsing request body", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	if strings.TrimSpace(strengthExercise.Name) == "" {
		logger.Logger.Error("error creating strength exercise, strength exercise name cannot be empty")
		c.JSON(http.StatusBadRequest, gin.H{"error": "exercise name is required"})
		return
	}

	strengthExercise.Name = strings.ToLower(strengthExercise.Name)
	err = sec.service.Create(ctx, &strengthExercise)
	if err != nil {
		logger.Logger.Error("error creating strength exercise,error saving strength exercise to database", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("error creating strength exercise, %s", err.Error())})
		return
	}
	logger.Logger.Info("successfully created strength exercise", "name", strengthExercise.Name)
	c.JSON(http.StatusCreated, gin.H{"data": strengthExercise})
}

// PUT method to edit a Strength Exercise record
func (sec *strengthExerciseController) Put(c *gin.Context) {
	ctx := c.Request.Context()
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		logger.Logger.Error("error updating strength exercise, string converstion to interger error", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid strength exercise ID"})
		return
	}
	logger.Logger.Info("recevied request to update strength exercise by id", "id", id)

	var input models.StrengthExercise
	if err := c.ShouldBindBodyWithJSON(&input); err != nil {
		logger.Logger.Error("error updating strength exercise, error parsing request body", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	if strings.TrimSpace(input.Name) == "" {
		logger.Logger.Error("error updating strength exercise, strength exercise name cannot be empty", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "strength exercise name is required"})
		return
	}

	input.Name = strings.ToLower(input.Name)
	err = sec.service.Update(ctx, uint(id), input)
	if err != nil {
		logger.Logger.Error("error updating strength exercise in database", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("error updating strength exercise %s", err.Error())})
		return
	}
	logger.Logger.Info("successfully updated strength exercise by id", "id", id)
	c.JSON(http.StatusOK, gin.H{"message": "strength exercise updated successfully"})
}

// DELETE method to delete an existing Strength Exercise records
func (sec *strengthExerciseController) Delete(c *gin.Context) {
	ctx := c.Request.Context()
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		logger.Logger.Error("error deleting strength exercise, string converstion to interger error", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid strength exercise ID"})
		return
	}
	logger.Logger.Info("received request to delete strength exwercise by id", "id", id)

	if err := sec.service.Delete(ctx, uint(id)); err != nil {
		logger.Logger.Error("error deleting strength exercise", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("error deleting strength exercise, %s", err.Error())})
		return
	}
	logger.Logger.Info("successfully deleted strength exercise by id", "id", id)
	c.JSON(http.StatusOK, gin.H{"message": "strength exercise deleted successfully"})
}
