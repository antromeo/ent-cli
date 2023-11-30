package app

import (
	"context"
	"fmt"
	"github.com/antromeo/ent-cli/v2/utilities"
	"github.com/spf13/cobra"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"os"
	"sort"
)

// configmapCmd represents the configmap command
var configmapCmd = &cobra.Command{
	Use:   "configmap",
	Short: "show the images configmap used to deploy the current EntandoApp",
	Long:  "show the images configmap used to deploy the current EntandoApp",
	Run: func(cmd *cobra.Command, args []string) {
		k8sClient := utilities.GetKubeClientInstance()
		configMapName := "entando-docker-image-info"
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
	AppCmd.AddCommand(configmapCmd)
}
