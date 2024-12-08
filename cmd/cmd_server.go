/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/AmitKarnam/WorkoutTracker/database/mysql"
	"github.com/AmitKarnam/WorkoutTracker/models"
	"github.com/AmitKarnam/WorkoutTracker/server"
	"github.com/joho/godotenv"
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

		//Start Server
		eg.Go(func() error {
			err := server.Start(port)
			if err != nil {
				return err
			}
			return nil
		})

		// Load environment variables
		if err = godotenv.Load(); err != nil {
			return fmt.Errorf("Error loading .env file")
		}

		// TODO : This needs to come from env variables
		dbPath := "amit:amit@tcp(localhost:3307)/workout_tracker?parseTime=true"

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
		fmt.Println("Completed migration database..")

		eg.Wait()

		return nil
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)

	serverCmd.Flags().StringVarP(&port, "port", "p", "9000", "--port 9000 or -p 9000")
}
