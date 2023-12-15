package app

import (
	"context"
	"fmt"
	"github.com/antromeo/ent-cli/v2/utilities"
	"github.com/spf13/cobra"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"os"
)

// editCmd represents the edit command
var editCmd = &cobra.Command{
	Use:   "edit [entandoAppName]",
	Short: "make changes in entando app",
	Long:  "make changes in entando app",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		replicas, _ := cmd.Flags().GetInt("replicas")
		k8sClient := utilities.GetKubeClientInstance()
		enap, err := k8sClient.EntandoClientSet.EntandoV1().EntandoApps(k8sClient.Namespace).Get(context.TODO(), args[0], metav1.GetOptions{})

		if err != nil {
			fmt.Printf("error get entando app: %v\n", err)
			os.Exit(1)
		}
		if replicas > 0 {
			enap.Spec.Replicas = replicas
		}

		// force redeployment
		enap.Annotations["entando.org/processing-instruction"] = "force"

		_, err = k8sClient.EntandoClientSet.EntandoV1().EntandoApps(k8sClient.Namespace).Update(context.TODO(), enap, metav1.UpdateOptions{})
		if err != nil {
			fmt.Printf("error updating entando app: %v\n", err)
			os.Exit(1)
		}

		fmt.Println("entando app updated!")
	},
}

func init() {
	AppCmd.AddCommand(editCmd)
	editCmd.Flags().Int("replicas", 0, "number of replica")
}
