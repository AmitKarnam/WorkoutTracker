package server

import (
	"github.com/AmitKarnam/WorkoutTracker/internal/controller"
	"github.com/AmitKarnam/WorkoutTracker/internal/controller/oauth/google"
	"github.com/AmitKarnam/WorkoutTracker/internal/middleware"
	"github.com/AmitKarnam/WorkoutTracker/internal/models"
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
				muscleGroup := versionGroup.Group("/muscle-groups")
				muscleGroupRepository := repository.NewMuscleGroupRepository(db)
				userRepository := repository.NewUserRepository(db)
				userService := services.NewUserService(userRepository)

				muscleGroupService := services.NewMuscleGroupService(muscleGroupRepository)

				muscleGroupController := controller.NewMuscleGroupController(muscleGroupService)
				muscleGroup.Use(middleware.JWTValidate(userService), middleware.RBAC(models.Admin))
				muscleGroup.GET("", muscleGroupController.Get)
				muscleGroup.GET(":id", muscleGroupController.GetByID)
				muscleGroup.POST("", muscleGroupController.Post)
				muscleGroup.PUT("/update/:id", muscleGroupController.Put)
				muscleGroup.DELETE("/delete/:id", muscleGroupController.Delete)
			}

			{
				strengthExercise := versionGroup.Group("strength-exercises")
				muscleGroupRepository := repository.NewMuscleGroupRepository(db)
				userRepository := repository.NewUserRepository(db)
				userService := services.NewUserService(userRepository)
				strengthExerciseRepository := repository.NewStrengthExerciseRepository(db, muscleGroupRepository)
				strengthExerciseService := services.NewStrengthExerciseService(strengthExerciseRepository)

				strengthExercise.Use(middleware.JWTValidate(userService), middleware.RBAC(models.Admin))

				strengthExerciseController := controller.NewStrengthExerciseController(strengthExerciseService)
				strengthExercise.GET("", strengthExerciseController.Get)
				strengthExercise.GET(":id", strengthExerciseController.GetByID)
				strengthExercise.POST("", strengthExerciseController.Post)
				strengthExercise.PUT("/update/:id", strengthExerciseController.Put)
				strengthExercise.DELETE("/delete/:id", strengthExerciseController.Delete)
			}

			{
				yogaExercise := versionGroup.Group("yoga-exercises")
				userRepository := repository.NewUserRepository(db)
				userService := services.NewUserService(userRepository)

				yogaExercise.Use(middleware.JWTValidate(userService), middleware.RBAC(models.Admin))

				yogaExercise.GET("")
				yogaExercise.GET(":id")
				yogaExercise.POST("")
				yogaExercise.PUT("/update/:id")
				yogaExercise.DELETE("/delete/:id")
			}

			{
				coreExercise := versionGroup.Group("core-exercises")
				userRepository := repository.NewUserRepository(db)
				userService := services.NewUserService(userRepository)

				coreExercise.Use(middleware.JWTValidate(userService), middleware.RBAC(models.Admin))

				coreExercise.GET("")
				coreExercise.GET(":id")
				coreExercise.POST("")
				coreExercise.PUT("/update/:id")
				coreExercise.DELETE("/delete/:id")
			}
		}
	}
}
