package cmd

import (
	"bytes"
	"context"
	"fmt"
	"github.com/antromeo/ent-cli/v2/constants"
	"github.com/antromeo/ent-cli/v2/utilities"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"os"
	"os/exec"
	"slices"
)

// bundleCmd represents the bundle command
var bundleCmd = &cobra.Command{
	Use:   "bundle",
	Short: "Management of new generation entando bundles",
	Long:  "Management of new generation entando bundles",
	Run: func(cmd *cobra.Command, args []string) {

		os.Setenv("ENTANDO_BUNDLE_CLI_BIN_NAME", "ent-cli bundle")

		if slices.Contains(args, "--debug") {
			os.Setenv("ENTANDO_CLI_DEBUG", "true")
			args = removeValue(args, "--debug")
		}

		if slices.Contains(args, "deploy") {
			deployOnCluster()
		} else {
			entandoConfig := utilities.GetEntandoConfigInstance()
			cmd := exec.Command(entandoConfig.GetEntBundleCliBinFilePath(), args...)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			_ = cmd.Run()
		}

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

func deployOnCluster() {
	entandoConfig := utilities.GetEntandoConfigInstance()
	cmd := exec.Command(entandoConfig.GetEntBundleCliBinFilePath(), "generate-cr")
	output, err := cmd.CombinedOutput()

	if err != nil {
		fmt.Println(string(output))
		os.Exit(1)
	}

	fmt.Println(string(output))
	cr := normalizeYaml(output)
	unstructuredCr := &unstructured.Unstructured{
		Object: map[string]interface{}{},
	}
	if err := yaml.Unmarshal(cr, &unstructuredCr.Object); err != nil {
		fmt.Printf("Error parsing YAML: %v\n", err)
		return
	}

	k8sClient := utilities.GetKubeClientInstance()

	_, err = k8sClient.DynamicClient.
		Resource(constants.EntandoDeBundleGroupVersionResource).
		Namespace(k8sClient.Namespace).
		Create(context.TODO(), unstructuredCr, metav1.CreateOptions{})

	if err != nil {
		fmt.Printf("Error deploying the bundle: %v\n", err)
		os.Exit(1)
	} else {
		fmt.Printf("Bundle successfully deployed\n")
	}
}

func normalizeYaml(obj []byte) []byte {
	index := bytes.Index(obj, []byte("---"))
	if index != -1 {
		obj = obj[index:]
	}
	return obj
}

func removeValue(args []string, value string) []string {
	for i := 0; i < len(args); i++ {
		if args[i] == value {
			return append(args[:i], args[i+1:]...)
		}
	}
	return args
}
