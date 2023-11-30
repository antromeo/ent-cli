package app

import (
	"github.com/spf13/cobra"
)

// AppCmd represents the app command
var AppCmd = &cobra.Command{
	Use:   "app",
	Short: "Helper",
	Long:  "Helper",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
}
