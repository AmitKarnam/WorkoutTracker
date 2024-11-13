package controllers

// import (
// 	"fmt"
// 	"net/http"

// 	"github.com/AmitKarnam/WorkoutTracker/models"

// 	"github.com/gin-gonic/gin"
// )

// type ExerciseController struct{}

// func (ec *ExerciseController) Get(c *gin.Context) {
// 	c.JSON(http.StatusOK, gin.H{"data": models.ExerciseList})
// }

// // this method is based on the fact that an existent muscle group will be passed as argument
// func (ec *ExerciseController) GetByMuscleGroup(c *gin.Context) {
// 	musclegroup := c.Param("musclegroup")
// 	var exerciseMuscleGroup []models.Exercise
// 	for _, exercise := range models.ExerciseList {
// 		if exercise.MuscleGroup == models.MuscleGroup(musclegroup) {
// 			exerciseMuscleGroup = append(exerciseMuscleGroup, exercise)
// 		}
// 	}
// 	if exerciseMuscleGroup != nil {
// 		c.JSON(http.StatusOK, gin.H{"data": exerciseMuscleGroup})
// 	} else {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "Exercise for the specified muscle group not found"})
// 	}
// }

// func (ec *ExerciseController) Post(c *gin.Context) {
// 	var newExercise models.Exercise
// 	if err := c.ShouldBindJSON(&newExercise); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 	}
// 	models.ExerciseList = append(models.ExerciseList, newExercise)
// 	c.JSON(http.StatusCreated, gin.H{"data": newExercise})
// }

// func (ec *ExerciseController) Put(c *gin.Context) {
// 	name := c.Param("name")
// 	for i, exercise := range models.ExerciseList {
// 		if exercise.Name == name {
// 			if err := c.ShouldBindJSON(&exercise); err != nil {
// 				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			}
// 			models.ExerciseList[i] = exercise
// 			c.JSON(http.StatusOK, gin.H{"data": exercise})
// 		}
// 	}
// 	c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Exercise with name %s not found", name)})
// }

// func (ec *ExerciseController) Delete(c *gin.Context) {
// 	name := c.Param("name")
// 	for i, exercise := range models.ExerciseList {
// 		if exercise.Name == name {
// 			models.ExerciseList = append(models.ExerciseList[:i], models.ExerciseList[i+1:]...)
// 			c.JSON(http.StatusNoContent, gin.H{})
// 		}
// 	}
// 	c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Exercise with name %s not found", name)})
// }
