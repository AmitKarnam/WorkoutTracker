/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/AmitKarnam/WorkoutTracker/server"

	"github.com/spf13/cobra"
)

var port string

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		// Start Server
		err := server.Start(port)
		if err != nil {
			return err
		}

		// Load environment variables

		return nil
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)

	serverCmd.Flags().StringVarP(&port, "port", "p", "9000", "--port 9000 or -p 9000")
}
