package profile

import (
	"fmt"
	"github.com/antromeo/ent-cli/v2/utilities"
	"github.com/spf13/cobra"
)

// currentCmd represents the current command
var currentCmd = &cobra.Command{
	Use:   "current",
	Short: "Shows the current profile",
	Long:  "Shows the current profiles",
	Run: func(cmd *cobra.Command, args []string) {
		entandoConfig := utilities.GetEntandoConfigInstance()
		fmt.Printf("Currently using profile \"%s\"\n", entandoConfig.GetProfile())
	},
}

func init() {
	ProfileCmd.AddCommand(currentCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// currentCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// currentCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
