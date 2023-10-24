package profile

import (
	"fmt"
	"github.com/antromeo/ent-cli/v2/constants"
	"github.com/antromeo/ent-cli/v2/utilities"
	"github.com/spf13/cobra"
	"os"
	"slices"
)

// useCmd represents the set command
var useCmd = &cobra.Command{
	Use:        "use [profileName]",
	Short:      "Selects the profile that ent should use",
	Long:       "Selects the profile that ent should use",
	Args:       cobra.ExactArgs(1),
	ArgAliases: []string{"profileName"},

	Run: func(cmd *cobra.Command, args []string) {
		entandoConfig := utilities.GetEntandoConfigInstance()
		profile := args[0]
		profiles, _ := entandoConfig.GetProfiles()

		if !slices.Contains(profiles, profile) {
			fmt.Println("Profile selected not found")
			os.Exit(1)
		}

		var config constants.Config
		utilities.ReadFileToYaml(entandoConfig.GetEntGlobalConfigFilePath(), &config)

		config.DesignedProfile = profile

		utilities.WriteYamlToFile(entandoConfig.GetEntGlobalConfigFilePath(), config)

		fmt.Println("The designated profile has been updated")

	},
}

func init() {
	ProfileCmd.AddCommand(useCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// setCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// setCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
