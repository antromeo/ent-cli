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
		fmt.Printf("Currently using profile \"%s\"\n", utilities.GetProfile())
		files, err := utilities.GetProfiles()
		if err != nil {
			return
		}
		fmt.Println("profile availables:")
		for _, file := range files {
			fmt.Println(file.Name())
		}
	},
}

func init() {
	ProfileCmd.AddCommand(listCmd)
}
