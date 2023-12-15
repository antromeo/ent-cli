package ecr

import (
	"fmt"
	"github.com/antromeo/ent-cli/v2/utilities"
	"github.com/spf13/cobra"
)

// getTenantIdCmd represents the getTenantId command
var getTenantIdCmd = &cobra.Command{
	Use:   "get-tenant-id",
	Short: "calculate tenant id",
	Long:  "calculate tenant id",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("%v\n", utilities.HashAndTruncate(args[0]))
	},
}

func init() {
	EcrCmd.AddCommand(getTenantIdCmd)
}
