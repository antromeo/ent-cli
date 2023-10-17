package profile

import (
	"ent-cli/constants"
	"ent-cli/utilities"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

// setCmd represents the set command
var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Sets appname and namespace of the current profile",
	Long:  "sets appname and namespace of the current profile",
	Run: func(cmd *cobra.Command, args []string) {
		appName, _ := cmd.Flags().GetString("appname")
		namespace, _ := cmd.Flags().GetString("namespace")

		configFilePath := utilities.GetEntConfigFilePathByProfile(utilities.GetProfile())
		var config constants.ProfileConfig
		utilities.ReadFileToYaml(configFilePath, &config)

		if utilities.IsEmpty(appName) {
			appName = utilities.ReadString("Please provide the EntandoApp name", false)
		}

		if utilities.IsEmpty(namespace) {
			namespace = utilities.ReadString("Please provide the Namespace", false)
		}

		if utilities.IsEmpty(appName) && utilities.IsEmpty(namespace) {
			fmt.Println("Nothing to update")
			os.Exit(1)
		}

		if !utilities.IsEmpty(appName) {
			config.AppName = appName
		}

		if !utilities.IsEmpty(namespace) {
			config.Namespace = namespace
		}

		utilities.WriteYamlToFile(configFilePath, config)
		fmt.Println("Profile updated")

	},
}

func init() {
	ProfileCmd.AddCommand(setCmd)

	setCmd.Flags().String("appname", "", "Entando AppName")
	setCmd.Flags().String("namespace", "", "Namespace")
}
