package config

import (
	"fmt"
	"github.com/antromeo/ent-cli/v2/utilities"
	"os"

	"github.com/spf13/cobra"
)

// setCmd represents the set command
var setCmd = &cobra.Command{
	Use:   "set",
	Short: "sets a specific configuration parameter",
	Long:  "sets a specific configuration parameter",
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		if (len(args) % 2) != 0 {
			fmt.Printf("requires an even number of args, received %v\n", len(args))
			os.Exit(1)
		}
		entandoConfig := utilities.GetEntandoConfigInstance()
		configFilePath := entandoConfig.GetEntConfigFilePathByProfile(entandoConfig.GetProfile())
		var config map[string]string
		utilities.ReadFileToYaml(configFilePath, &config)
		for i := 0; i < len(args); i += 2 {
			key := args[i]
			value := args[i+1]
			config[key] = value
		}
		utilities.WriteYamlToFile(configFilePath, config)
		fmt.Println("The properties have been updated")

	},
}

func init() {
	ConfigCmd.AddCommand(setCmd)
}
