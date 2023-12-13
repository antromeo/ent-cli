package ecr

import (
	"github.com/antromeo/ent-cli/v2/digitalexchange"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"os"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list bundles and show their status",
	Long:  "list bundles and show their status",
	Run: func(cmd *cobra.Command, args []string) {
		ecrComponentsResponse := digitalexchange.GetEcrComponents()

		if len(ecrComponentsResponse.Payload) > 0 {
			table := tablewriter.NewWriter(os.Stdout)
			table.SetHeader([]string{"Bundle Code", "Status", "Version"})

			for _, bundle := range ecrComponentsResponse.Payload {
				table.Append([]string{bundle.Code, bundle.LastJob.Status, bundle.LastJob.ComponentVersion})
			}

			table.Render()
		}

	},
}

func init() {
	EcrCmd.AddCommand(listCmd)
}
