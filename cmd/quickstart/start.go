package quickstart

import (
	"fmt"
	"github.com/antromeo/ent-cli/v2/constants"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "start a Entando local instance",
	Long:  "start a Entando local instance",
	Run: func(cmd *cobra.Command, args []string) {

		name, _ := cmd.Flags().GetString("name")
		cpu, _ := cmd.Flags().GetString("cpu")
		memory, _ := cmd.Flags().GetString("memory")
		kubeVersion, _ := cmd.Flags().GetString("kubernetes-version")
		namespace, _ := cmd.Flags().GetString("namespace")
		entandoVersion, _ := cmd.Flags().GetString("entando-version")
		enableTekton, _ := cmd.Flags().GetBool("enable-tekton")

		cmdLineArgs := []string{
			"start", "-p", name,
			"--cpus", cpu, "--memory", memory,
			"--driver", "docker",
			"--container-runtime", "cri-o",
			"--addons", "ingress,default-storageclass,storage-provisioner",
			"--kubernetes-version", kubeVersion,
		}

		execCmd := exec.Command("minikube", cmdLineArgs[0:]...)
		execCmd.Stdout = os.Stdout
		execCmd.Stderr = os.Stderr

		err := execCmd.Run()
		if err != nil {
			fmt.Printf("Error creating local instance: %v\n", err)
			os.Exit(1)
		}

		fmt.Println("instance created")

		// setting profile
		execCmd = exec.Command("minikube", "profile", name)
		err = execCmd.Run()
		if err != nil {
			fmt.Printf("Error setting profile: %v\n", err)
			os.Exit(1)
		}

		execCmd = exec.Command("minikube", "ip", "-p", name)
		output, err := execCmd.CombinedOutput()
		if err != nil {
			fmt.Printf("Error getting IP from local instance: %v\n", err)
			os.Exit(1)
		}
		instanceIp := strings.Replace(string(output), "\n", "", -1)

		fmt.Printf("instance available at address: %s\n", instanceIp)

		execCmd = exec.Command("minikube", "kubectl", "create", "namespace", namespace)
		err = execCmd.Run()
		if err != nil {
			fmt.Printf("Error creating namespace: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("namespace %s created\n", namespace)

		applyEntandoTemplates(namespace, entandoVersion)
		applyEntandoApp(namespace, instanceIp)

		appBuilderAddress := strings.Join([]string{"quickstart", instanceIp, "nip.io/app-builder/"}, ".")
		fmt.Printf("entando resources applied, you can find the Entando Application at: %s\n", appBuilderAddress)

		if enableTekton {
			tektonTemplateUrl := "https://storage.googleapis.com/tekton-releases/pipeline/latest/release.yaml"
			execCmd := exec.Command("minikube", "kubectl", "--", "apply", "-f", tektonTemplateUrl)
			err = execCmd.Run()
			if err != nil {
				fmt.Printf("Error applied the tekton templates: %v\n", err)
				os.Exit(1)
			}
		}

	},
}

func init() {
	QuickstartCmd.AddCommand(startCmd)
	startCmd.Flags().String("name", "", "Name of profile")
	startCmd.Flags().String("cpu", "8", "Number of CPUs allocated to Kubernetes. Use \"max\" to use the maximum number of CPUs.\n")
	startCmd.Flags().String("memory", "16G", "Amount of RAM to allocate to Kubernetes (format: <number>[<unit>], where unit = b, k, m or g). Use \"max\" to\n\tuse the maximum amount of memory.")
	startCmd.Flags().String("kubernetes-version", "v1.26.9", "kubernetes version")
	startCmd.Flags().String("entando-version", "", "entando version")
	startCmd.Flags().String("namespace", "entando", "namespace to use")
	startCmd.Flags().Bool("enable-tekton", false, "enable tekton in cluster")
	startCmd.MarkFlagRequired("name")
	startCmd.MarkFlagRequired("entando-version")
}

func applyEntandoTemplates(namespace string, version string) {
	for _, urlTemplate := range constants.EntandoResourcesTemplates {
		templateUrl := fmt.Sprintf(urlTemplate, version)
		execCmd := exec.Command("minikube", "kubectl", "--", "-n", namespace, "apply", "-f", templateUrl)
		err := execCmd.Run()
		if err != nil {
			fmt.Printf("Error applied the templates: %v\n", err)
			os.Exit(1)
		}
	}
}

func applyEntandoApp(namespace string, ip string) {
	entandoApp := getEntandoApp(namespace, ip)
	entandoAppYaml, err := yaml.Marshal(entandoApp.Object)
	if err != nil {
		fmt.Printf("Error marshaling to YAML: %v\n", err)
		return
	}
	entandoAppFilePath := filepath.Join(os.TempDir(), "entandoApp.yaml")
	err = os.WriteFile(entandoAppFilePath, entandoAppYaml, 0600)
	if err != nil {
		fmt.Printf("Error writing template: %v\n", err)
		os.Exit(1)
	}
	execCmd := exec.Command("minikube", "kubectl", "--", "apply", "-f", entandoAppFilePath)
	err = execCmd.Run()
	if err != nil {
		fmt.Printf("Error applied EntandoApp: %v\n", err)
		os.Exit(1)
	}

}

func getEntandoApp(namespace string, ip string) *unstructured.Unstructured {
	hostname := strings.Join([]string{"quickstart", ip, "nip.io"}, ".")
	return &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "entando.org/v1",
			"kind":       "EntandoApp",
			"metadata": map[string]interface{}{
				"namespace": namespace,
				"name":      "quickstart",
			},
			"spec": map[string]interface{}{
				"environmentVariables": []map[string]interface{}{
					{
						"name":  "SPRING_PROFILES_ACTIVE",
						"value": "default,swagger",
					},
				},
				"dbms":                "postgresql",
				"ingressHostName":     hostname,
				"standardServerImage": "tomcat",
				"replicas":            1,
			},
		},
	}
}
