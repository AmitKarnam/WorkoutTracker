package server

import (
	"github.com/AmitKarnam/WorkoutTracker/internal/controller"
	"github.com/AmitKarnam/WorkoutTracker/internal/controller/oauth/google"
	"github.com/AmitKarnam/WorkoutTracker/internal/middleware"
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
				authGroup := versionGroup.Group("/auth")
				userRepo := repository.NewUserRepository(db)
				refreshTokenRepo := repository.NewRefreshTokenRepository(db)
				refreshTokenService := services.NewRefreshTokenService(refreshTokenRepo)

				authGroup.POST("/logout")

				tokenGroup := authGroup.Group("/token")
				{
					refreshTokenController := controller.NewRefreshTokenController(refreshTokenService)
					tokenGroup.POST("/refresh", refreshTokenController.RefreshTokenHandler)
				}

				googleGroup := authGroup.Group("/google")
				{
					googleAuthService := services.NewGoogleAuthService(userRepo)
					googleAuthController := google.NewGoogleAuthController(googleAuthService, refreshTokenService)
					googleGroup.GET("/login", googleAuthController.LoginHandler)
					googleGroup.GET("/callback", googleAuthController.CallbackHandler)
				}
			}

			{

				// TODO: Should be under authorization middleware
				// TODO: Should be under RBAC middleware
				muscleGroup := versionGroup.Group("/muscle-groups")
				muscleGroupRepository := repository.NewMuscleGroupRepository(db)
				userRepository := repository.NewUserRepository(db)
				userService := services.NewUserService(userRepository)
				muscleGroupService := services.NewMuscleGroupService(muscleGroupRepository)
				muscleGroupController := controller.NewMuscleGroupController(muscleGroupService)
				muscleGroup.Use(func(c *gin.Context) { middleware.JWTValidate(c, userService) })
				muscleGroup.GET("", muscleGroupController.Get)
				muscleGroup.GET(":id", muscleGroupController.GetByID)
				muscleGroup.POST("", muscleGroupController.Post)
				muscleGroup.PUT("/update/:id", muscleGroupController.Put)
				muscleGroup.DELETE("/delete/:id", muscleGroupController.Delete)
			}

			{
				// TODO: Should be under authorization middleware
				// TODO: Should be under RBAC middleware
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
				// TODO: Should be under authorization middleware
				// TODO: Should be under RBAC middleware
				yogaExercise := versionGroup.Group("yoga-exercises")
				yogaExercise.GET("")
				yogaExercise.GET(":id")
				yogaExercise.POST("")
				yogaExercise.PUT("/update/:id")
				yogaExercise.DELETE("/delete/:id")
			}

			{
				// TODO: Should be under authorization middleware
				// TODO: Should be under RBAC middleware
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
