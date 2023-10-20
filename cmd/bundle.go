package cmd

import (
	"ent-cli/constants"
	"ent-cli/utilities"
	"fmt"
	"os/exec"
	"slices"

	"github.com/spf13/cobra"
)

// bundleCmd represents the bundle command
var bundleCmd = &cobra.Command{
	Use:   "bundle",
	Short: "Management of new generation entando bundles",
	Long:  "Management of new generation entando bundles",
	Run: func(cmd *cobra.Command, args []string) {
		var versions map[string]interface{}
		var entBundleVersion = constants.DefaultTag
		utilities.ReadFileToYaml("releases/releases.yaml", &versions)
		if value, found := versions["entando_bundle_cli"]; found {
			fmt.Printf("entando_bundle_cli: %v\n", value)
			entBundleVersion = value.(string)
		} else {
			fmt.Println("entando_bundle_cli not found in the YAML data")
		}
		image := fmt.Sprint("docker.io/romeocontainers/entando-bundle-cli:", entBundleVersion)
		cmdLine := append([]string{"run", "-t", image, "bin/run"}, args...)

		if slices.Contains(args, "deploy") {
			// TODO: implementation
		}

		podmanCmd := exec.Command("podman", cmdLine[0:]...)

		// Capture the command's standard output and error
		output, _ := podmanCmd.CombinedOutput()

		// Print the output to the console
		fmt.Printf(string(output))

		if slices.Contains(args, "--help") {
			utilities.ShowAdditionalCommandsInHelp()
		}
		return
	},
}

func init() {
	bundleCmd.DisableFlagParsing = true
	rootCmd.AddCommand(bundleCmd)
}
