package profile

import (
	"ent-cli/utilities"
	"fmt"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Shows a list of available profiles",
	Long:  "Shows a list of available profiles",
	Run: func(cmd *cobra.Command, args []string) {
		entandoConfig := utilities.GetEntandoConfigInstance()
		profile := entandoConfig.GetProfile()
		fmt.Printf("Currently using profile \"%s\"\n", profile)
		files, err := entandoConfig.GetProfiles()
		if err != nil {
			return
		}
		fmt.Println("profile availables:")
		for _, file := range files {
			fmt.Println(file)
		}
	},
}

func init() {
	ProfileCmd.AddCommand(listCmd)
}
