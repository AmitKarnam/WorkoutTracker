package server

import (
	"github.com/AmitKarnam/WorkoutTracker/internal/controllers"
	"github.com/AmitKarnam/WorkoutTracker/internal/repository"
	"github.com/AmitKarnam/WorkoutTracker/internal/services"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

func initRoutes(engine *gin.Engine, db *gorm.DB) {
	serviceGroup := engine.Group("workouttracker")
	{
		// Health endpoint
		health := serviceGroup.Group("/health")
		healthController := controllers.HealthController{}
		health.GET("", healthController.Get)

		apiGroup := serviceGroup.Group("/api")
		{
			versionGroup := apiGroup.Group("/v1")

			{
				exerciseCategory := versionGroup.Group("/exercisecategories")
				exerciseCategoryRepository := repository.NewExerciseCategoryRepository(db)
				exerciseCategoryServices := services.NewExerciseCategoryService(exerciseCategoryRepository)
				exerciseCategoryController := controllers.NewExerciseCategoryController(exerciseCategoryServices)
				exerciseCategory.GET("", exerciseCategoryController.Get)
				exerciseCategory.GET(":id", exerciseCategoryController.GetByID)
				exerciseCategory.POST("", exerciseCategoryController.Post)
				exerciseCategory.DELETE("/delete/:id", exerciseCategoryController.Delete)
			}

			{
				muscleGroup := versionGroup.Group("/musclegroups")
				muscleGroupRepository := repository.NewMuscleGroupRepository(db)
				muscleGroupService := services.NewMuscleGroupService(muscleGroupRepository)
				muscleGroupController := controllers.NewMuscleGroupController(muscleGroupService)
				muscleGroup.GET("", muscleGroupController.Get)
				muscleGroup.GET(":id", muscleGroupController.GetByID)
				muscleGroup.POST("", muscleGroupController.Post)
				muscleGroup.PUT("/update/:id", muscleGroupController.Put)
				muscleGroup.DELETE("/delete/:id", muscleGroupController.Delete)
			}
		}
	}
}
