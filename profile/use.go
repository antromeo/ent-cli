package profile

import (
	"ent-cli/constants"
	"ent-cli/utilities"
	"fmt"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
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

		fileContent, err := os.ReadFile(entandoConfig.GetEntGlobalConfigFilePath())
		if err != nil {
			fmt.Println("Error loading the profile", err)
			return
		}
		var config constants.Config
		if err := yaml.Unmarshal(fileContent, &config); err != nil {
			fmt.Println("Error unmarshaling profile config:", err)
			return
		}

		config.DesignedProfile = profile

		updatedContent, err := yaml.Marshal(config)
		if err != nil {
			fmt.Println("Error marshaling profile config:", err)
			return
		}

		if err := os.WriteFile(entandoConfig.GetEntGlobalConfigFilePath(), updatedContent, 0600); err != nil {
			fmt.Println("Error writing to the file:", err)
			return
		}
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
