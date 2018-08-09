package cmd

import (
	"github.com/amy911/srv911/run"

	"github.com/spf13/cobra"
)

// Set this in an `init()` function somewhere
var OverrideListen func(*cobra.Command, []string)

// listenCmd represents the listen command
var listenCmd = &cobra.Command{
	Use:   "listen",
	Short: "Listen and serve",
	Long: `Listen and serve`,
	Run: func(cmd *cobra.Command, args []string) {
		switch {
		case OverrideListen == nil:
			run.Listen(cmd, args)
		default:
			OverrideListen(cmd, args)
		}
	},
}

func init() {
	rootCmd.AddCommand(listenCmd)
}
