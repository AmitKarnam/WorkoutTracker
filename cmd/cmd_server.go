/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/AmitKarnam/WorkoutTracker/database/mysql"
	"github.com/AmitKarnam/WorkoutTracker/models"
	"github.com/AmitKarnam/WorkoutTracker/server"
	"github.com/joho/godotenv"
	"github.com/knadh/koanf"
	"github.com/knadh/koanf/providers/env"
	"github.com/spf13/cobra"
	"golang.org/x/sync/errgroup"
)

var port string

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		var eg errgroup.Group
		var err error

		// Load environment variables from .env file (optional)
		if err := godotenv.Load(); err != nil {
			log.Println("No .env file found, proceeding with environment variables")
		}

		// Create a new koanf instance
		k := koanf.New(".")

		// Load environment variables into koanf
		if err := k.Load(env.Provider("", ".", func(s string) string {
			return s // No transformation on environment variable keys
		}), nil); err != nil {
			log.Fatalf("Error loading environment variables: %v", err)
		}

		// Read values from koanf
		user := k.String("MYSQL_DATABASE_USER")
		password := k.String("MYSQL_DATABASE_PASSWORD")
		host := k.String("MYSQL_HOST")
		mysql_port := k.String("MYSQL_PORT")
		database := k.String("MYSQL_DATABASE")

		// Create the database connection string
		dbPath := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, password, host, mysql_port, database)

		//Start Server
		eg.Go(func() error {
			err := server.Start(port)
			if err != nil {
				return err
			}
			return nil
		})

		// Initialise DB
		mysql.NewMySQLInit(dbPath)

		// Get database connection and migrate the models
		dbConn, err := mysql.DB.GetConnection()
		if err != nil {
			return err
		}
		dbConn.AutoMigrate(&models.MuscleGroup{})
		dbConn.AutoMigrate(&models.Exercise{})
		dbConn.AutoMigrate(&models.Workout{})
		log.Println("Completed migration database...")

		log.Println("Server initialised...")

		eg.Wait()

		return nil
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)

	serverCmd.Flags().StringVarP(&port, "port", "p", "9000", "--port 9000 or -p 9000")
}
