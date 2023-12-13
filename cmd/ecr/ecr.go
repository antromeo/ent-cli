package ecr

import (
	"github.com/spf13/cobra"
)

// EcrCmd represents the ecr command
var EcrCmd = &cobra.Command{
	Use:   "ecr",
	Short: "Helper for managing the ECR",
	Long:  "Helper for managing the ECR",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
}
