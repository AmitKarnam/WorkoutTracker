package controller

import (
	"github.com/AmitKarnam/WorkoutTracker/internal/services"
	"github.com/gin-gonic/gin"
)

type StrengthExerciseController interface {
	Get(*gin.Context)
	GetByID(*gin.Context)
	Post(*gin.Context)
	Put(*gin.Context)
	Delete(*gin.Context)
}

type strengthExerciseController struct {
	service services.StrengthExerciseService
}

func NewStrengthExerciseController(service services.StrengthExerciseService) StrengthExerciseController {
	return &strengthExerciseController{service: service}
}

// Get All Strength Exercises from database
func (sec *strengthExerciseController) Get(ctx *gin.Context) {

}

// Get sepcific Strength Exercise by ID from database
func (sec *strengthExerciseController) GetByID(ctx *gin.Context) {

}

// Post method to add a new Strength Exercise to database
func (sec *strengthExerciseController) Post(ctx *gin.Context) {

}

// PUT method to edit a Strength Exercise record
func (sec *strengthExerciseController) Put(ctx *gin.Context) {

}

// DELETE method to delete an existing Strength Exercise records
func (sec *strengthExerciseController) Delete(ctx *gin.Context) {

}
