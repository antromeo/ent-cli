package ecr

import (
	"fmt"
	"github.com/antromeo/ent-cli/v2/utilities"
	"github.com/spf13/cobra"
	"strings"
)

// getPluginCodeCmd represents the getPluginCode command
var getPluginCodeCmd = &cobra.Command{
	Use:   "get-plugin-code",
	Short: "calculates and displays the plugin code",
	Long:  "calculates and displays the plugin code",
	Run: func(cmd *cobra.Command, args []string) {
		msName, _ := cmd.Flags().GetString("ms-name")
		bundleName, _ := cmd.Flags().GetString("bundle-name")
		org, _ := cmd.Flags().GetString("org")
		registry, _ := cmd.Flags().GetString("registry")
		tenant, _ := cmd.Flags().GetString("tenant")

		if utilities.IsEmpty(msName) {
			msName = utilities.ReadString("Please provide the microservice name", true)
		}

		if utilities.IsEmpty(bundleName) {
			bundleName = utilities.ReadString("Please provide the bundle name", true)
		}

		if utilities.IsEmpty(org) {
			org = utilities.ReadString("Please provide the organization name", true)
		}

		bundleFqdnRepo := strings.Join([]string{registry, org, bundleName}, "/")
		msNameWithOrg := strings.Join([]string{org, msName}, "/")

		pluginCode := strings.Join([]string{"pn", utilities.HashAndTruncate(bundleFqdnRepo), utilities.HashAndTruncate(msNameWithOrg), utilities.NormalizeName(msNameWithOrg)}, "-")

		if tenant != "primary" {
			tenantCode := utilities.HashAndTruncate(tenant)
			pluginCode = strings.Join([]string{"pn", utilities.HashAndTruncate(bundleFqdnRepo), tenantCode, utilities.HashAndTruncate(msNameWithOrg), utilities.NormalizeName(msNameWithOrg)}, "-")
		}

		fmt.Printf("%v\n", pluginCode)

	},
}

func init() {
	EcrCmd.AddCommand(getPluginCodeCmd)
	getPluginCodeCmd.Flags().String("ms-name", "", "microservice name")
	getPluginCodeCmd.Flags().String("bundle-name", "", "bundle name")
	getPluginCodeCmd.Flags().String("org", "", "your organization")
	getPluginCodeCmd.Flags().String("registry", "registry.hub.docker.com", "your registry")
	getPluginCodeCmd.Flags().String("tenant", "primary", "tenant name")
}
