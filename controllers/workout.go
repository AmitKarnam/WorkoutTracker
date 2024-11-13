package controllers

// import (
// 	"net/http"

// 	"github.com/AmitKarnam/WorkoutTracker/models"

// 	"github.com/gin-gonic/gin"
// )

// type WorkoutController struct{}

// func (wc *WorkoutController) Get(c *gin.Context) {

// }

// func (wc *WorkoutController) GetByName(c *gin.Context) {
// 	exerciseName := c.Param("exercise")
// 	var newExerciseList []models.Exercise

// 	for _, exercise := range models.ExerciseList {
// 		if exercise.Name == exerciseName {
// 			newExerciseList = append(newExerciseList, exercise)
// 		}
// 	}

// 	c.JSON(http.StatusAccepted, gin.H{"data": newExerciseList})
// }

// func (wc *WorkoutController) GetWorkoutByMuscleGroup(c *gin.Context) {

// }

// func (wc *WorkoutController) Post(c *gin.Context) {
// 	var newWorkoutEntry models.Workout

// 	if err := c.ShouldBindBodyWithJSON(&newWorkoutEntry); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 	}
// 	models.WorkoutList = append(models.WorkoutList, newWorkoutEntry)
// 	c.JSON(http.StatusAccepted, gin.H{"data": newWorkoutEntry})
// }
