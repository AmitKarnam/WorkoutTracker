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

type ExerciseCategoryController interface {
	Get(*gin.Context)
	GetByID(*gin.Context)
	Post(*gin.Context)
	Delete(*gin.Context)
}

type exerciseCategoryController struct {
	service services.ExerciseCategoryService
}

func NewExerciseCategoryController(service services.ExerciseCategoryService) ExerciseCategoryController {
	return &exerciseCategoryController{service: service}
}

// Get method to get all exercise categories
func (e *exerciseCategoryController) Get(c *gin.Context) {
	logger.Logger.Info("received request to get all exercise categories")
	exerciseCategories, err := e.service.GetAll()
	if err != nil {
		logger.Logger.Error("error fetching exercise categories", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error fetching exercise categories"})
		return
	}
	logger.Logger.Info("successfully fetched all exercise categories")
	c.JSON(http.StatusOK, gin.H{"data": exerciseCategories})
}

// GetByID methods gets exercise category by id
func (e *exerciseCategoryController) GetByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		logger.Logger.Error("error fetching exercise category by id, string conversion to integer error", "id", id)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid exercise category ID"})
		return
	}
	logger.Logger.Info("received request to fetch exercise category by id", "id", id)

	exerciseCategory, err := e.service.GetByID(uint(id))
	if err != nil {
		logger.Logger.Error("error fetching exercise group by id", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("error fetching exercise category by id, %s", err.Error())})
		return
	}
	logger.Logger.Info("successfully fetched exercise category by id", "id", id)
	c.JSON(http.StatusOK, gin.H{"data": exerciseCategory})
}

// Post method is used to create exercise category
func (e *exerciseCategoryController) Post(c *gin.Context) {
	logger.Logger.Info("received request to create new exercise category")
	var exerciseCategory models.ExerciseCategory
	if err := c.ShouldBindJSON(&exerciseCategory); err != nil {
		logger.Logger.Error("error creating exercise category, error parsing request body", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	if strings.TrimSpace(exerciseCategory.Category) == "" {
		logger.Logger.Error("error creating exercise category, exercise category cannot be empty")
		c.JSON(http.StatusBadRequest, gin.H{"error": "exercise category is required"})
		return
	}

	exerciseCategory.Category = strings.ToLower(exerciseCategory.Category)

	err := e.service.Create(&exerciseCategory)
	if err != nil {
		logger.Logger.Error("error creating exercise category, error saving exercise category to database", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("error creating exercise category, %s", err.Error())})
		return
	}
	logger.Logger.Info("successfully created exercise category", "name", exerciseCategory.Category)
	c.JSON(http.StatusCreated, gin.H{"data": exerciseCategory})
}

func (e *exerciseCategoryController) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		logger.Logger.Error("error deleting exercise, string converstion to integer error", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("invalid exercise category ID")})
		return
	}
	logger.Logger.Info("received request to delete exercise category by id", "id", id)

	if err := e.service.Delete(uint(id)); err != nil {
		logger.Logger.Error("error deleting exercise category", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("error deleting exercise category, %s", err.Error())})
		return
	}

	logger.Logger.Info("successfully deleted exercise by id", "id", id)
	c.JSON(http.StatusOK, gin.H{"message": "exercise category deleted successfully"})
}
