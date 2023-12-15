package ecr

import (
	"fmt"
	"github.com/antromeo/ent-cli/v2/utilities"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

// getBundleIdCmd represents the getBundleId command
var getBundleIdCmd = &cobra.Command{
	Use:   "get-bundle-id",
	Short: "calculates and displays the bundle id",
	Long:  "calculates and displays the bundle id",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if !utilities.IsValidURL(args[0]) {
			fmt.Println("The URL is not valid.")
			os.Exit(1)
		}
		// Extract the registry from the repository URL
		repoUrl := strings.SplitN(args[0], "://", 2)
		repoName := repoUrl[1]
		if len(repoUrl) != 2 {
			fmt.Printf("invalid input format\n")
			os.Exit(1)
		}
		fmt.Printf("%v\n", utilities.HashAndTruncate(repoName))
	},
}

func init() {
	EcrCmd.AddCommand(getBundleIdCmd)
}
