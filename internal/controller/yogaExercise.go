package controller

import (
	"github.com/AmitKarnam/WorkoutTracker/internal/services"
	"github.com/gin-gonic/gin"
)

type YogaExerciseController interface {
	Get(*gin.Context)
	GetByID(*gin.Context)
	Post(*gin.Context)
	Put(*gin.Context)
	Delete(*gin.Context)
}

type yogaExerciseController struct {
	service services.StrengthExerciseService
}

func NewYogaExerciseController(service services.StrengthExerciseService) YogaExerciseController {
	return &yogaExerciseController{service: service}
}

// Get All Strength Exercises from database
func (sec *yogaExerciseController) Get(ctx *gin.Context) {

}

// Get sepcific Strength Exercise by ID from database
func (sec *yogaExerciseController) GetByID(ctx *gin.Context) {

}

// Post method to add a new Strength Exercise to database
func (sec *yogaExerciseController) Post(ctx *gin.Context) {

}

// PUT method to edit a Strength Exercise record
func (sec *yogaExerciseController) Put(ctx *gin.Context) {

}

// DELETE method to delete an existing Strength Exercise records
func (sec *yogaExerciseController) Delete(ctx *gin.Context) {

}
