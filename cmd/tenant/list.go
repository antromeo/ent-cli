package tenant

import (
	"context"
	"fmt"
	"github.com/antromeo/ent-cli/v2/utilities"
	"github.com/spf13/cobra"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/json"
)

// listCmd represents the tenantlist command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "shows list of tenants",
	Long:  "shows list of tenants",
	Run: func(cmd *cobra.Command, args []string) {
		k8sClient := utilities.GetKubeClientInstance()
		tenantConfigSecret, _ := k8sClient.ClientSet.CoreV1().Secrets(k8sClient.Namespace).Get(context.TODO(), "tenant-config-secret", metav1.GetOptions{})
		tenantConfigBytes := tenantConfigSecret.Data["TENANT_CONFIGS"]

		var tenantConfig []entandoConfigs
		err := json.Unmarshal(tenantConfigBytes, &tenantConfig)
		if err != nil {
			fmt.Println("Error decoding JSON:", err)
			return
		}
		for _, value := range tenantConfig {
			fmt.Printf("%s\n", value.TenantCode)
		}
	},
}

func init() {
	TenantCmd.AddCommand(listCmd)
}

type entandoConfigs struct {
	TenantCode string `json:"tenantCode"`
}
