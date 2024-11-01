/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/AmitKarnam/WorkoutTracker/models"
	"github.com/AmitKarnam/WorkoutTracker/server"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"golang.org/x/sync/errgroup"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var port string

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		var eg errgroup.Group

		//Start Server
		eg.Go(func() error {
			err := server.Start(port)
			if err != nil {
				return err
			}
			return nil
		})

		// Load environment variables
		if err := godotenv.Load(); err != nil {
			return fmt.Errorf("Error loading .env file")
		}

		// Connecting to Database
		fmt.Printf("Connecting to database and migrating database.....")
		dsn := "amit:amit@tcp(localhost:3307)/workout_tracker"
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			fmt.Errorf("Error connectin to database : ", err)
		}

		db.AutoMigrate(&models.Exercise{})
		//db.AutoMigrate(&models.Workout{})

		eg.Wait()

		return nil
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)

	serverCmd.Flags().StringVarP(&port, "port", "p", "9000", "--port 9000 or -p 9000")
}
