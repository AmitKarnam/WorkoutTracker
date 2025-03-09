package initialiseServer

import (
	"fmt"

	"github.com/AmitKarnam/WorkoutTracker/database/mysql"
	"github.com/AmitKarnam/WorkoutTracker/internal/models"
	"github.com/AmitKarnam/WorkoutTracker/internal/server"
	"github.com/AmitKarnam/WorkoutTracker/logger"
	"github.com/joho/godotenv"
	"github.com/knadh/koanf"
	"github.com/knadh/koanf/providers/env"
	"golang.org/x/sync/errgroup"
)

var k *koanf.Koanf

func init() {
	k = koanf.New(".")
}

// InitServer funciton is starts up all the essential services required for the server to run
func InitServer() error {
	var eg errgroup.Group
	// Initialise logger
	initLogger()

	err := loadEnv("./.env")
	if err != nil {
		logger.Logger.Error("Failed to load environment variables", "error", err)
		return err
	}
	logger.Logger.Info("Loaded environment variables successfully")

	err = migrateDB()
	if err != nil {
		logger.Logger.Error("Failed to migrate database", "error", err)
		return err
	}
	logger.Logger.Info("Migrated database successfully")

	eg.Go(func() error {
		err := startRESTServer()
		if err != nil {
			logger.Logger.Error("Failed to start REST server", "error", err)
			return err
		}
		return nil
	})
	fmt.Println("Started REST server successfully......")
	logger.Logger.Info("Started REST server successfully")
	eg.Wait()
	return nil
}

// Initialise logger
func initLogger() {
	logger.InitLogger()
}

// Load environemnt variables
func loadEnv(filepath string) error {
	// Load environment variables
	if err := godotenv.Load(filepath); err != nil {
		logger.Logger.Info("No .env file found, proceeding with environment variables")
	}

	// Load environment variables into koanf
	if err := k.Load(env.Provider("", ".", func(s string) string {
		return s
	}), nil); err != nil {
		return err
	}
	return nil
}

// Load secrets
// func loadSecrets() error {
// 	return nil
// }

// Connect to database and migrate database
func migrateDB() error {
	// Read values from koanf
	user := k.String("DATABASE_USER")
	password := k.String("DATABASE_PASSWORD")
	host := k.String("DATABASE_HOST")
	mysql_port := k.String("DATABASE_PORT")
	database := k.String("DATABASE")

	// Create the database connection string
	dbPath := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, password, host, mysql_port, database)

	// Initialise DB
	mysql.NewMySQLInit(dbPath)

	// Get database connection and migrate the models
	dbConn, err := mysql.DB.GetConnection()
	if err != nil {
		return err
	}

	logger.Logger.Info("Starting database migration...")

	err = dbConn.AutoMigrate(&models.ExerciseCategory{}, &models.MuscleGroup{}, &models.StrengthExercise{}, &models.StrengthWorkout{}, &models.CoreExercise{}, &models.CoreWorkout{}, &models.YogaExercise{}, &models.YogaWorkout{})
	if err != nil {
		return err
	}

	logger.Logger.Info("Completed migration database...")

	return nil
}

// Start server
func startRESTServer() error {
	var serverPort string
	serverPort = k.String("SERVER_PORT")
	err := server.Start(serverPort)
	if err != nil {
		return err
	}
	return nil
}
