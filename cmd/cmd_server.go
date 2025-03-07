package cmd

import (
	"github.com/AmitKarnam/WorkoutTracker/internal/initialiseServer"
	"github.com/AmitKarnam/WorkoutTracker/logger"
	"github.com/spf13/cobra"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{

	Use:   "server",
	Short: "",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := initialiseServer.InitServer()
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
}
