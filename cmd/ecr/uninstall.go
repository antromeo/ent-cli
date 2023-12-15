package ecr

import (
	"fmt"
	"github.com/antromeo/ent-cli/v2/digitalexchange"
	"github.com/spf13/cobra"
)

// uninstallCmd represents the uninstall command
var uninstallCmd = &cobra.Command{
	Use:   "uninstall [bundle-code]",
	Short: "uninstalls a bundle",
	Long:  "uninstalls a bundle",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		res := digitalexchange.UninstallComponent(args[0])
		fmt.Printf("request sent with status: %v\n", res.Payload.Status)
	},
}

func init() {
	EcrCmd.AddCommand(uninstallCmd)
}
