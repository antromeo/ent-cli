package profile

import (
	"github.com/spf13/cobra"
)

// ProfileCmd represents the profile command
var ProfileCmd = &cobra.Command{
	Use:   "profile",
	Short: "Handles ent profiles",
	Long:  `Handles ent profiles`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
}
