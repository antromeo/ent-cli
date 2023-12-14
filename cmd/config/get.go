package config

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "gets a specific configuration parameter",
	Long:  "gets a specific configuration parameter",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		value := viper.GetString(args[0])
		if len(value) > 0 {
			fmt.Printf("property found, key: \"%v\", value: \"%v\"\n", args[0], value)
		} else {
			fmt.Printf("property with key: \"%v\" not found\n", args[0])
		}
	},
}

func init() {
	ConfigCmd.AddCommand(getCmd)
}
