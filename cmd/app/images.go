package app

import (
	"context"
	"fmt"
	"github.com/antromeo/ent-cli/v2/utilities"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"os"
	"sort"
	"strings"
)

// imagesCmd represents the images command
var imagesCmd = &cobra.Command{
	Use:   "images",
	Short: "show images info",
	Long:  "show images info",
	Run: func(cmd *cobra.Command, args []string) {
		k8sClient := utilities.GetKubeClientInstance()

		pods, err := k8sClient.ClientSet.CoreV1().Pods(k8sClient.Namespace).List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			fmt.Printf("Error listing pods: %v\n", err)
			os.Exit(1)
		}

		sort.SliceStable(pods.Items, func(i, j int) bool {
			return strings.Compare(pods.Items[i].ObjectMeta.Name, pods.Items[j].ObjectMeta.Name) < 0
		})

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Pod Name", "Container Name", "Container Image"})

		for _, pod := range pods.Items {
			for _, container := range pod.Spec.Containers {
				table.Append([]string{pod.ObjectMeta.Name, container.Name, container.Image})
			}
		}

		table.Render()
	},
}

func init() {
	AppCmd.AddCommand(imagesCmd)
}
