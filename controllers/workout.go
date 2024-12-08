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

type WorkoutController struct{}

// GetWorkouts method to fetch workouts based on query parameters
func (wc *WorkoutController) GetWorkouts(c *gin.Context) {

	date := c.DefaultQuery("date", "")                // Single date (optional)
	startDate := c.DefaultQuery("start_date", "")     // Start date for date range (optional)
	endDate := c.DefaultQuery("end_date", "")         // End date for date range (optional)
	muscleGroup := c.DefaultQuery("muscle_group", "") // Muscle group name (optional)
	exerciseName := c.DefaultQuery("exercise", "")    // Exercise name (optional)

	dbConn, err := mysql.DB.GetConnection()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error connecting to database"})
		return
	}

	var workouts []models.Workout
	query := dbConn.Preload("Exercise")

	if date != "" {
		query = query.Where("DATE(created_at) = ?", date)
	}

	if startDate != "" && endDate != "" {
		query = query.Where("created_at BETWEEN ? AND ?", startDate, endDate)
	}

	if muscleGroup != "" {
		query = query.Joins("JOIN exercises ON exercises.id = workouts.exercise_id").
			Joins("JOIN muscle_groups ON muscle_groups.id = exercises.muscle_group_id").
			Where("muscle_groups.muscle_group = ?", muscleGroup)
	}

	if exerciseName != "" {
		query = query.Where("exercises.name = ?", exerciseName)
	}

	if err := query.Find(&workouts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error fetching workouts"})
		return
	}

	if len(workouts) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "no workouts found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": workouts})
}

// Post method to add a new workout record
func (wc *WorkoutController) Post(c *gin.Context) {
	var workout models.Workout

	if err := c.ShouldBindJSON(&workout); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	dbConn, err := mysql.DB.GetConnection()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error connecting to database"})
		return
	}

	var exercise models.Exercise
	if err := dbConn.First(&exercise, workout.ExerciseID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "exercise not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error fetching exercise"})
		return
	}

	if err := dbConn.Create(&workout).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error saving workout"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": workout})
}

// Put updates a workout by its ID
func (wc *WorkoutController) Put(c *gin.Context) {
	workoutID := c.Param("id")

	id, err := strconv.ParseUint(workoutID, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid workout ID"})
		return
	}

	dbConn, err := mysql.DB.GetConnection()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error connecting to database"})
		return
	}

	var workout models.Workout
	if err := dbConn.First(&workout, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "workout not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error fetching workout"})
		return
	}

	var input models.Workout
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	workout.ExerciseID = input.ExerciseID
	workout.Reps = input.Reps
	workout.Weight = input.Weight

	if err := dbConn.Save(&workout).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error updating workout"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": workout})
}

// Delete method to remove a workout by ID
func (wc *WorkoutController) Delete(c *gin.Context) {
	id := c.Param("id")

	dbConn, err := mysql.DB.GetConnection()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error connecting to database"})
		return
	}

	var workout models.Workout
	if err := dbConn.First(&workout, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "workout not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error fetching workout"})
		return
	}

	if err := dbConn.Delete(&workout).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error deleting workout"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "workout deleted successfully"})
}
