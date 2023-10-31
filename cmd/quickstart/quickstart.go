package quickstart

import (
	"github.com/spf13/cobra"
)

// quickstartCmd represents the quickstart command
var QuickstartCmd = &cobra.Command{
	Use:   "quickstart",
	Short: "Helper for installing Entando instances locally",
	Long:  "Helper for installing Entando instances locally",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
}
