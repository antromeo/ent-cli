package config

import (
	"github.com/spf13/cobra"
)

// configCmd represents the config command
var ConfigCmd = &cobra.Command{
	Use:   "config",
	Short: "",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {

}
