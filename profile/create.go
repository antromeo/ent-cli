package profile

import (
	. "ent-cli/constants"
	"ent-cli/utilities"
	"fmt"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"path/filepath"
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
		home, _ := os.UserHomeDir()

		// Marshal the configuration to YAML format.
		yamlData, err := yaml.Marshal(&profileConfig)
		if err != nil {
			log.Fatalf("Error marshaling YAML: %v", err)
		}
		profileFolder := filepath.Join(home, EntFolder, ProfilesFolder, name)
		err = os.MkdirAll(profileFolder, os.ModePerm)
		if err != nil {
			fmt.Printf("Error creating directory: %v\n", err)
			return
		}
		err = os.WriteFile(filepath.Join(profileFolder, ConfigFile+".yaml"), yamlData, 0644)
		if err != nil {
			log.Fatalf("Error writing to file: %v", err)
		}
		fmt.Printf("Profile created\n")
	},
}

func init() {
	ProfileCmd.AddCommand(createCmd)

	createCmd.Flags().String("name", "", "Profile Name")
	createCmd.Flags().String("appname", "", "Entando AppName")
	createCmd.Flags().String("namespace", "", "Namespace")
}
