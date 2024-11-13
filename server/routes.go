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

			// {
			// 	muscleGroup := versionGroup.Group("/musclegroups")
			// 	muscleGroupController := controllers.MuscleGroupController{}
			// 	muscleGroup.GET("", muscleGroupController.Get)
			// 	muscleGroup.POST("", muscleGroupController.Post)
			// 	//muscleGroup.DELETE(":id", muscleGroupController.Delete)
			// }

			// {
			// 	exerciseGroup := versionGroup.Group("/exercises")
			// 	exerciseController := controllers.ExerciseController{}
			// 	exerciseGroup.GET("", exerciseController.Get)
			// 	exerciseGroup.GET(":musclegroup", exerciseController.GetByMuscleGroup)
			// 	exerciseGroup.POST("", exerciseController.Post)
			// 	exerciseGroup.PUT(":id", exerciseController.Put)
			// 	exerciseGroup.DELETE(":id", exerciseController.Delete)
			// }

			{
				workoutGroup := versionGroup.Group("/workouts")
				workoutGroup.GET("")
				workoutGroup.POST("")
				workoutGroup.PUT(":id")
				workoutGroup.DELETE(":id")
			}
		}
	}
}
