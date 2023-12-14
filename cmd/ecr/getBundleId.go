package ecr

import (
	"crypto/sha256"
	"encoding/hex"
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
		fmt.Printf("%v\n", hashAndTruncate(args[0]))
	},
}

func hashAndTruncate(input string) string {
	// Extract the registry from the repository URL
	repoUrl := strings.SplitN(input, "://", 2)
	if len(repoUrl) != 2 {
		fmt.Printf("invalid input format\n")
		os.Exit(1)
	}
	repoName := repoUrl[1]
	hash := sha256.New()
	hash.Write([]byte(repoName))
	hashString := hex.EncodeToString(hash.Sum(nil))

	return hashString[:8]

}

func init() {
	EcrCmd.AddCommand(getBundleIdCmd)
}
