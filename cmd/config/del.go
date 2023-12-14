package config

import (
	"fmt"
	"github.com/antromeo/ent-cli/v2/utilities"
	"github.com/spf13/cobra"
)

// delCmd represents the del command
var delCmd = &cobra.Command{
	Use:   "del",
	Short: "deletes a specific configuration parameter",
	Long:  "deletes a specific configuration parameter",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		entandoConfig := utilities.GetEntandoConfigInstance()
		configFilePath := entandoConfig.GetEntConfigFilePathByProfile(entandoConfig.GetProfile())
		var config map[string]string
		utilities.ReadFileToYaml(configFilePath, &config)
		for _, arg := range args {
			if _, ok := config[arg]; ok {
				delete(config, arg)
			} else {
				fmt.Println("Env", arg, "not found.")
			}
		}
		utilities.WriteYamlToFile(configFilePath, config)
		fmt.Println("The properties have been updated")
	},
}

func init() {
	ConfigCmd.AddCommand(delCmd)
}
