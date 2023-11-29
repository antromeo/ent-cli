package cmd

import (
	"context"
	"fmt"
	"github.com/antromeo/ent-cli/v2/utilities"
	"github.com/spf13/cobra"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// statusCmd represents the status command
var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Status informations",
	Long:  "Status informations",
	Run: func(cmd *cobra.Command, args []string) {
		k8sClient := utilities.GetKubeClientInstance()

		debundleList, _ := k8sClient.EntandoClientSet.EntandoV1().EntandoDeBundles(k8sClient.Namespace).List(context.TODO(), metav1.ListOptions{})
		fmt.Println("List of Entando Bundles")
		for _, debundle := range debundleList.Items {
			fmt.Println(debundle.Name)
		}
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)

}
