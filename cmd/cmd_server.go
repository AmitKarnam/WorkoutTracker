/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	initialiseserver "github.com/AmitKarnam/WorkoutTracker/internal/initaliseserver"
	"github.com/AmitKarnam/WorkoutTracker/logger"
	"github.com/spf13/cobra"
)

var port string

// serverCmd represents the server command
var serverCmd = &cobra.Command{

	Use:   "server",
	Short: "",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := initialiseserver.InitServer(port)
		if err != nil {
			logger.Logger.Error("Failed to initialise server", "error", err)
			return err
		}
		logger.Logger.Info("Initialised wourkout server successfully")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)

	serverCmd.Flags().StringVarP(&port, "port", "p", "9000", "--port <port> or -p <port>")
}
