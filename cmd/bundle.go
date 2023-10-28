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
	"strings"
)

// bundleCmd represents the bundle command
var bundleCmd = &cobra.Command{
	Use:   "bundle",
	Short: "Management of new generation entando bundles",
	Long:  "Management of new generation entando bundles",
	Run: func(cmd *cobra.Command, args []string) {
		entBundleImage := strings.Join([]string{constants.EntandoBundleRepository, constants.EntandoBundleCliVersion}, ":")
		cmdLine := append([]string{"run", "-t", entBundleImage}, args...)

		if slices.Contains(args, "deploy") {
			deployOnCluster(entBundleImage)
		} else {
			podmanCmd := exec.Command(constants.ContainerRuntime, cmdLine[0:]...)
			output, _ := podmanCmd.CombinedOutput()
			fmt.Printf(string(output))
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

func deployOnCluster(entBundleImage string) {
	cmdLine := append([]string{"run", "-t", "-v", ".:/app", entBundleImage}, "generate-cr")
	podmanCmd := exec.Command(constants.ContainerRuntime, cmdLine[0:]...)
	output, err := podmanCmd.CombinedOutput()

	if err != nil {
		fmt.Println(string(output))
		os.Exit(1)
	}

	fmt.Println(string(output))
	cr := normalizeYaml(output)
	unstructedCr := &unstructured.Unstructured{
		Object: map[string]interface{}{},
	}
	if err := yaml.Unmarshal(cr, &unstructedCr.Object); err != nil {
		fmt.Printf("Error parsing YAML: %v\n", err)
		return
	}

	k8sClient := utilities.GetKubeClientInstance()

	_, err = k8sClient.DynamicClient.
		Resource(constants.EntandoDeBundleGroupVersionResource).
		Namespace(k8sClient.Namespace).
		Create(context.TODO(), unstructedCr, metav1.CreateOptions{})

	if err != nil {
		fmt.Printf("Error deploying the bundle: %v\n", err)
		os.Exit(1)
	} else {
		fmt.Printf("Bundle successfully deployed")
	}
}

func normalizeYaml(obj []byte) []byte {
	index := bytes.Index(obj, []byte("---"))
	if index != -1 {
		obj = obj[index:]
	}
	return obj
}
