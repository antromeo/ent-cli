package quickstart

import (
	"github.com/spf13/cobra"
	"os"
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
		execCmd.Stdout = os.Stdout
		execCmd.Stderr = os.Stderr
		_ = execCmd.Run()
	},
}

func init() {
	QuickstartCmd.AddCommand(deleteCmd)
	deleteCmd.Flags().String("name", "", "Name of profile")
	deleteCmd.Flags().Bool("all", false, "Set flag to delete all local instances")
	deleteCmd.MarkFlagRequired("name")
}
