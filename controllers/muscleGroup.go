package controllers

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/AmitKarnam/WorkoutTracker/database/mysql"
	"github.com/AmitKarnam/WorkoutTracker/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type MuscleGroupController struct{}

// Get method to fetch all muscle groups from database
func (msc *MuscleGroupController) Get(c *gin.Context) {

	name := c.Query("name")

	dbConn, err := mysql.DB.GetConnection()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error connecting to database"})
		return
	}

	var muscleGroups []models.MuscleGroup
	if name != "" {
		err = dbConn.Where("muscle_group = ?", name).Find(&muscleGroups).Error
	} else {
		err = dbConn.Find(&muscleGroups).Error
	}

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

// Put method to edit a muscle group record
func (mc *MuscleGroupController) Put(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.ParseUint(idStr, 10, 32) // Convert string ID to uint
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid muscle group ID"})
		return
	}

	dbConn, err := mysql.DB.GetConnection()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error connecting to database"})
		return
	}

	var muscleGroup models.MuscleGroup
	if err := dbConn.First(&muscleGroup, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "muscle group not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error fetching muscle group"})
		return
	}

	var input models.MuscleGroup
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	// Update fields from input
	muscleGroup.MuscleGroup = input.MuscleGroup
	muscleGroup.Description = input.Description

	if err := dbConn.Save(&muscleGroup).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error updating muscle group"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": muscleGroup})
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
