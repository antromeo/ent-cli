package quickstart

import (
	"fmt"
	"github.com/spf13/cobra"
	"os/exec"
)

// stopCmd represents the stop command
var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "stop local entando instance",
	Long:  "stop local entando instance",
	Run: func(cmd *cobra.Command, args []string) {
		all, _ := cmd.Flags().GetBool("all")
		var cmdLineArgs []string

		if all {
			cmdLineArgs = []string{"stop", "--all"}
		} else {
			name, _ := cmd.Flags().GetString("name")
			cmdLineArgs = []string{"stop", "-p", name}
		}

		execCmd := exec.Command("minikube", cmdLineArgs[0:]...)
		output, _ := execCmd.CombinedOutput()
		fmt.Printf(string(output))
	},
}

func init() {
	QuickstartCmd.AddCommand(stopCmd)
	stopCmd.Flags().String("name", "", "Name of profile")
	stopCmd.Flags().Bool("all", false, "Set flag to stop all local instances")
	stopCmd.MarkFlagRequired("name")
}
