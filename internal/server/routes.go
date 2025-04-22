package server

import (
	"github.com/AmitKarnam/WorkoutTracker/internal/controller"
	"github.com/AmitKarnam/WorkoutTracker/internal/repository"
	"github.com/AmitKarnam/WorkoutTracker/internal/services"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

func initRoutes(engine *gin.Engine, db *gorm.DB) {
	serviceGroup := engine.Group("workout-tracker")
	{
		// Health endpoint
		health := serviceGroup.Group("/health")
		healthController := controller.HealthController{}
		health.GET("", healthController.Get)

		apiGroup := serviceGroup.Group("/api")
		{
			versionGroup := apiGroup.Group("/v1")

			{
				authRepo := repository.NewUserRepository(db)
				authService := services.NewAuthService(authRepo)
				authController := controller.NewAuthController(authService)

				authGroup := versionGroup.Group("/auth")
				authGroup.GET("/google/login", authController.GoogleLogin)
				authGroup.GET("/google/callback", authController.GoogleCallback)
			}

			{
				muscleGroup := versionGroup.Group("/muscle-groups")
				muscleGroupRepository := repository.NewMuscleGroupRepository(db)
				muscleGroupService := services.NewMuscleGroupService(muscleGroupRepository)
				muscleGroupController := controller.NewMuscleGroupController(muscleGroupService)
				muscleGroup.GET("", muscleGroupController.Get)
				muscleGroup.GET(":id", muscleGroupController.GetByID)
				muscleGroup.POST("", muscleGroupController.Post)
				muscleGroup.PUT("/update/:id", muscleGroupController.Put)
				muscleGroup.DELETE("/delete/:id", muscleGroupController.Delete)
			}

			{
				strengthExercise := versionGroup.Group("strength-exercises")
				muscleGroupRepository := repository.NewMuscleGroupRepository(db)
				strengthExerciseRepository := repository.NewStrengthExerciseRepository(db, muscleGroupRepository)
				strengthExerciseService := services.NewStrengthExerciseService(strengthExerciseRepository)
				strengthExerciseController := controller.NewStrengthExerciseController(strengthExerciseService)
				strengthExercise.GET("", strengthExerciseController.Get)
				strengthExercise.GET(":id", strengthExerciseController.GetByID)
				strengthExercise.POST("", strengthExerciseController.Post)
				strengthExercise.PUT("/update/:id", strengthExerciseController.Put)
				strengthExercise.DELETE("/delete/:id", strengthExerciseController.Delete)
			}

			{
				yogaExercise := versionGroup.Group("yoga-exercises")
				yogaExercise.GET("")
				yogaExercise.GET(":id")
				yogaExercise.POST("")
				yogaExercise.PUT("/update/:id")
				yogaExercise.DELETE("/delete/:id")
			}

			{
				coreExercise := versionGroup.Group("core-exercises")
				coreExercise.GET("")
				coreExercise.GET(":id")
				coreExercise.POST("")
				coreExercise.PUT("/update/:id")
				coreExercise.DELETE("/delete/:id")
			}
		}
	}
}
