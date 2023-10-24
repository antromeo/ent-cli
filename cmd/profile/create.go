package profile

import (
	"fmt"
	. "github.com/antromeo/ent-cli/v2/constants"
	"github.com/antromeo/ent-cli/v2/utilities"
	"github.com/spf13/cobra"
	"os"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new profile",
	Long:  "Create a new profile",
	Run: func(cmd *cobra.Command, args []string) {

		name, _ := cmd.Flags().GetString("name")
		appName, _ := cmd.Flags().GetString("appname")
		namespace, _ := cmd.Flags().GetString("namespace")

		if utilities.IsEmpty(name) {
			name = utilities.ReadString("Enter Profile Name", true)
		}
		if utilities.IsEmpty(appName) {
			appName = utilities.ReadString("Enter Entando App Name", true)
		}
		if utilities.IsEmpty(namespace) {
			namespace = utilities.ReadString("Enter Namespace", true)
		}

		profileConfig := ProfileConfig{
			AppName:   appName,
			Namespace: namespace,
		}

		entandoConfig := utilities.GetEntandoConfigInstance()
		profileFolder := entandoConfig.GetProfileFilePath(name)

		err := os.MkdirAll(profileFolder, 0770)
		if err != nil {
			fmt.Printf("Error creating directory: %v\n", err)
			return
		}

		utilities.WriteYamlToFile(entandoConfig.GetEntConfigFilePathByProfile(name), profileConfig)
		fmt.Printf("Profile created\n")
	},
}

func init() {
	ProfileCmd.AddCommand(createCmd)

	createCmd.Flags().String("name", "", "Profile Name")
	createCmd.Flags().String("appname", "", "Entando AppName")
	createCmd.Flags().String("namespace", "", "Namespace")
}
