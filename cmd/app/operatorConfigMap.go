package app

import (
	"context"
	"fmt"
	"github.com/antromeo/ent-cli/v2/utilities"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"os"
	"sort"

	"github.com/spf13/cobra"
)

// operatorConfigMapCmd represents the operatorConfigMap command
var operatorConfigMapCmd = &cobra.Command{
	Use:   "operator-configmap",
	Short: "show the configmap used configure the operator deployment parameters",
	Long:  "show the configmap used configure the operator deployment parameters",
	Run: func(cmd *cobra.Command, args []string) {
		k8sClient := utilities.GetKubeClientInstance()
		configMapName := "entando-operator-config"

		configMap, err := k8sClient.ClientSet.CoreV1().ConfigMaps(k8sClient.Namespace).Get(context.TODO(), configMapName, metav1.GetOptions{})
		if err != nil {
			fmt.Printf("Error getting ConfigMap: %v\n", err)
			os.Exit(1)
		}

		// Get the keys from the Data map and sort them alphabetically
		keys := make([]string, 0, len(configMap.Data))
		for key := range configMap.Data {
			keys = append(keys, key)
		}
		sort.Strings(keys)

		for _, key := range keys {
			value := configMap.Data[key]
			fmt.Printf("%s: '%s'\n", key, value)
		}

	},
}

func init() {
	AppCmd.AddCommand(operatorConfigMapCmd)
}
