package config

import (
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Show the configuration parameters",
	Long:  "Show the configuration parameters",
	Run: func(cmd *cobra.Command, args []string) {

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Env Name", "Value"})

		for _, env := range viper.AllKeys() {
			table.Append([]string{env, viper.GetString(env)})
		}

		table.Render()
	},
}

func init() {
	ConfigCmd.AddCommand(listCmd)

}
