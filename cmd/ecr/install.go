package ecr

import (
	"fmt"
	"github.com/antromeo/ent-cli/v2/digitalexchange"
	"github.com/antromeo/ent-cli/v2/utilities"
	"github.com/spf13/cobra"
	"os"
	"slices"
)

// installCmd represents the install command
var installCmd = &cobra.Command{
	Use:   "install",
	Short: "installs a bundle",
	Long:  "installs a bundle",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		strategy, err := cmd.Flags().GetString("strategy")
		if err != nil {
			fmt.Printf("Error retrieving strategy flag: %v\n", err)
			os.Exit(1)
		}

		validStrategies := []string{"CREATE", "SKIP", "OVERRIDE"}
		if !slices.Contains(validStrategies, strategy) {
			fmt.Printf("Invalid strategy: %s. Valid strategies are: %v\n", strategy, validStrategies)
			os.Exit(1)
		}

		version, err := cmd.Flags().GetString("version")
		if err != nil {
			fmt.Printf("Error retrieving version flag: %v\n", err)
			os.Exit(1)
		}
		if utilities.IsEmpty(version) {
			version = utilities.ReadString("Please provide the version to install", true)
		}

		res := digitalexchange.InstallComponent(args[0], version, strategy)
		fmt.Printf("request sent with status: %v\n", res.Payload.Status)

	},
}

func init() {
	EcrCmd.AddCommand(installCmd)
	installCmd.Flags().String("version", "", "semantic version of bundle")
	installCmd.Flags().String("strategy", "CREATE", "[CREATE|SKIP|OVERRIDE]")
}
