package profile

import (
	"fmt"
	"github.com/antromeo/ent-cli/v2/utilities"
	"github.com/spf13/cobra"
	"os"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a profile",
	Long:  "Delete a profile",
	Run: func(cmd *cobra.Command, args []string) {
		entandoConfig := utilities.GetEntandoConfigInstance()
		listProfiles, _ := entandoConfig.GetProfiles()
		for i, fileName := range listProfiles {
			fmt.Printf("%d) %s\n", i+1, fileName)
		}

		fmt.Print("Enter the number of the profile you want to delete: ")
		var selection int
		_, err := fmt.Scanln(&selection)
		if err != nil || selection < 1 || selection > len(listProfiles) {
			fmt.Println("Invalid selection.")
			return
		}

		selectedProfile := listProfiles[selection-1]

		err = os.RemoveAll(entandoConfig.GetProfileFilePath(selectedProfile))
		if err != nil {
			fmt.Printf("Error removing profile: %v\n", err)
			return
		}
		fmt.Printf("Profile %s removed\n", selectedProfile)

	},
}

func init() {
	ProfileCmd.AddCommand(deleteCmd)
}
