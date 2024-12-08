package controllers

import (
	"log"
	"net/http"

	"github.com/AmitKarnam/WorkoutTracker/database/mysql"
	"github.com/AmitKarnam/WorkoutTracker/models"
	"github.com/gin-gonic/gin"
)

type MuscleGroupController struct{}

// Get method to fetch all muscle groups from database
func (msc *MuscleGroupController) Get(c *gin.Context) {
	dbConn, err := mysql.DB.GetConnection()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error connecting to database"})
		return
	}

	var muscleGroups []models.MuscleGroup
	err = dbConn.Find(&muscleGroups).Error
	if err != nil {
		log.Println("Error fetching muscle groups:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error fetching muscle groups"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": muscleGroups})
}

// Post method to add a new muscle group to database
func (msc *MuscleGroupController) Post(c *gin.Context) {
	var muscleGroup models.MuscleGroup

	// Parse the JSON request body into muscleGroup
	if err := c.ShouldBindJSON(&muscleGroup); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Get database connection
	dbConn, err := mysql.DB.GetConnection()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error connecting to database"})
		return
	}

	// Check to if the value is already present, if so do not add the value
	var existingMuscleGroup models.MuscleGroup
	if err := dbConn.Where("muscle_group = ?", muscleGroup.MuscleGroup).First(&existingMuscleGroup).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "muscle group already exists"})
		return
	}

	// Save the muscle group to the database
	if err := dbConn.Create(&muscleGroup).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error saving muscle group"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": muscleGroup})
}

// Delete method to delete an existing muscle group from database
func (msc *MuscleGroupController) Delete(c *gin.Context) {
	muscleGroupName := c.Param("name") // Get the muscle group name from the URL parameters

	// Get database connection
	dbConn, err := mysql.DB.GetConnection()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error connecting to database"})
		return
	}

	// Find the muscle group by name
	var muscleGroup models.MuscleGroup
	if err := dbConn.Where("muscle_group = ?", muscleGroupName).First(&muscleGroup).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "muscle group not found"})
		return
	}

	// Delete the muscle group from the database
	if err := dbConn.Unscoped().Delete(&muscleGroup).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error deleting muscle group"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "muscle group deleted successfully"})
}
