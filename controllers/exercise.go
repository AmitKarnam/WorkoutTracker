package controllers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/AmitKarnam/WorkoutTracker/database/mysql"
	"github.com/AmitKarnam/WorkoutTracker/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ExerciseController struct{}

// GetExercises method to fetch exercises based on muscle group or exercise name query parameters
func (ec *ExerciseController) GetExercises(c *gin.Context) {
	muscleGroup := c.Query("muscle_group")
	exerciseName := c.Query("name")

	dbConn, err := mysql.DB.GetConnection()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error connecting to database"})
		return
	}

	var exercises []models.Exercise

	query := dbConn.Preload("MuscleGroup")

	if muscleGroup != "" {
		query = query.Joins("JOIN muscle_groups ON muscle_groups.id = exercises.muscle_group_id").
			Where("muscle_groups.muscle_group = ?", muscleGroup)
	}

	if exerciseName != "" {
		query = query.Where("exercises.name = ?", exerciseName)
	}

	if err := query.Find(&exercises).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error fetching exercises"})
		return
	}

	if len(exercises) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "no exercises found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": exercises})
}

// Post method to add new exercise record
func (ec *ExerciseController) Post(c *gin.Context) {
	var exercise models.Exercise

	if err := c.ShouldBindJSON(&exercise); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	dbConn, err := mysql.DB.GetConnection()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error connecting to database"})
		return
	}

	var existingExercise models.Exercise
	err = dbConn.Where("name = ?", exercise.Name).First(&existingExercise).Error
	if err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "exercise with the same name already exists"})
		return
	}

	if err := dbConn.Create(&exercise).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error saving exercise"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": exercise})
}

// UpdateByID method to update an exercise based on the ID in the URL parameter
func (ec *ExerciseController) Put(c *gin.Context) {
	idStr := c.Param("id") // Get the ID from the URL parameters

	id, err := strconv.ParseUint(idStr, 10, 32) // Convert string ID to uint
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid exercise ID"})
		return
	}

	var input models.Exercise
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	dbConn, err := mysql.DB.GetConnection()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error connecting to database"})
		return
	}

	var existingExercise models.Exercise
	if err := dbConn.First(&existingExercise, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "exercise not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error fetching exercise"})
		return
	}

	// Update fields from input
	existingExercise.Name = input.Name
	existingExercise.Description = input.Description
	existingExercise.MuscleGroupID = input.MuscleGroupID

	if err := dbConn.Save(&existingExercise).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error updating exercise"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": existingExercise})
}

// Delete method to remove an exercise by name
func (ec *ExerciseController) Delete(c *gin.Context) {
	name := c.Param("name") // Get the exercise name from the URL parameters

	dbConn, err := mysql.DB.GetConnection()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error connecting to database"})
		return
	}

	var exercise models.Exercise
	if err := dbConn.Where("name = ?", name).First(&exercise).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "exercise not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error fetching exercise"})
		return
	}

	if err := dbConn.Delete(&exercise).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error deleting exercise"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "exercise deleted successfully"})
}
