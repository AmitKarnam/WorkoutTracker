package controller

import (
	"github.com/AmitKarnam/WorkoutTracker/internal/services"
	"github.com/gin-gonic/gin"
)

type CoreExerciseController interface {
	Get(*gin.Context)
	GetByID(*gin.Context)
	Post(*gin.Context)
	Put(*gin.Context)
	Delete(*gin.Context)
}

type coreExerciseController struct {
	service services.StrengthExerciseService
}

func NewCoreExerciseController(service services.StrengthExerciseService) CoreExerciseController {
	return &coreExerciseController{service: service}
}

// Get All Strength Exercises from database
func (sec *coreExerciseController) Get(ctx *gin.Context) {

}

// Get sepcific Strength Exercise by ID from database
func (sec *coreExerciseController) GetByID(ctx *gin.Context) {

}

// Post method to add a new Strength Exercise to database
func (sec *coreExerciseController) Post(ctx *gin.Context) {

}

// PUT method to edit a Strength Exercise record
func (sec *coreExerciseController) Put(ctx *gin.Context) {

}

// DELETE method to delete an existing Strength Exercise records
func (sec *coreExerciseController) Delete(ctx *gin.Context) {

}
