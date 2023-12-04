package tenant

import (
	"github.com/spf13/cobra"
)

// TenantCmd represents the tenant command
var TenantCmd = &cobra.Command{
	Use:   "tenant",
	Short: "tenant informations",
	Long:  "tenant informations",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
}
