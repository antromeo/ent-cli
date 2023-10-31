package quickstart

import (
	"fmt"
	"github.com/spf13/cobra"
	"os/exec"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete local entando instance",
	Long:  "delete local entando instance",
	Run: func(cmd *cobra.Command, args []string) {
		all, _ := cmd.Flags().GetBool("all")
		var cmdLineArgs []string

		if all {
			cmdLineArgs = []string{"delete", "--all"}
		} else {
			name, _ := cmd.Flags().GetString("name")
			cmdLineArgs = []string{"delete", "-p", name}
		}

		execCmd := exec.Command("minikube", cmdLineArgs[0:]...)
		output, _ := execCmd.CombinedOutput()
		fmt.Printf(string(output))
	},
}

func init() {
	QuickstartCmd.AddCommand(deleteCmd)
	deleteCmd.Flags().String("name", "", "Name of profile")
	deleteCmd.Flags().Bool("all", false, "Set flag to delete all local instances")
	deleteCmd.MarkFlagRequired("name")
}
