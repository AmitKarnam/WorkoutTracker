package server

import (
	"github.com/AmitKarnam/WorkoutTracker/controllers"

	"github.com/gin-gonic/gin"
)

func initRoutes(engine *gin.Engine) {
	serviceGroup := engine.Group("wokouttracker")
	{
		// Health endpoint
		health := serviceGroup.Group("/health")
		healthController := controllers.HealthController{}
		health.GET("", healthController.Get)

		apiGroup := serviceGroup.Group("/api")
		{
			versionGroup := apiGroup.Group("/v1")

			{
				muscleGroup := versionGroup.Group("/musclegroups")
				muscleGroupController := controllers.MuscleGroupController{}
				muscleGroup.GET("", muscleGroupController.Get)
				muscleGroup.POST("", muscleGroupController.Post)
				muscleGroup.DELETE(":name", muscleGroupController.Delete)
			}

			{
				exerciseGroup := versionGroup.Group("/exercises")
				exerciseController := controllers.ExerciseController{}
				exerciseGroup.GET("", exerciseController.GetExercises)
				exerciseGroup.POST("", exerciseController.Post)
				exerciseGroup.PUT("", exerciseController.Put)
				exerciseGroup.DELETE(":name", exerciseController.Delete)
			}

			{
				workoutGroup := versionGroup.Group("/workouts")
				workoutController := controllers.WorkoutController{}
				workoutGroup.GET("", workoutController.GetWorkouts)
				workoutGroup.POST("", workoutController.Post)
				workoutGroup.PUT(":id", workoutController.Put)
				workoutGroup.DELETE(":id", workoutController.Delete)
			}
		}
	}
}
