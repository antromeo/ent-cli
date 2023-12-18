package app

import (
	"context"
	"fmt"
	"github.com/antromeo/ent-cli/v2/utilities"
	"github.com/spf13/cobra"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"os"
	"slices"
)

// editCmd represents the edit command
var editCmd = &cobra.Command{
	Use:   "edit [entandoAppName]",
	Short: "make changes in entando app",
	Long:  "make changes in entando app",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		replicas, _ := cmd.Flags().GetInt("replicas")
		server, _ := cmd.Flags().GetString("server")
		dbms, _ := cmd.Flags().GetString("dbms")
		k8sClient := utilities.GetKubeClientInstance()
		enap, err := k8sClient.EntandoClientSet.EntandoV1().EntandoApps(k8sClient.Namespace).Get(context.TODO(), args[0], metav1.GetOptions{})

		if err != nil {
			fmt.Printf("error get entando app: %v\n", err)
			os.Exit(1)
		}

		if replicas > 0 {
			enap.Spec.Replicas = replicas
		}

		if len(server) > 0 {
			validServer := []string{"tomcat", "wildfly", "eap"}
			if !slices.Contains(validServer, server) {
				fmt.Printf("Invalid server base image: %s. Valid values are: %v\n", server, validServer)
				os.Exit(1)
			}
			enap.Spec.StandardServerImage = server
		}

		if len(dbms) > 0 {
			validDbms := []string{"mysql", "oracle", "postgresql", "embedded", "none"}
			if !slices.Contains(validDbms, dbms) {
				fmt.Printf("Invalid dbms: %s. Valid values are: %v\n", dbms, validDbms)
				os.Exit(1)
			}
			enap.Spec.Dbms = dbms
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
	editCmd.Flags().String("server", "", "the standard server base image, allowed values: [tomcat|wildfly|eap]")
	editCmd.Flags().String("dbms", "", "dbms to use for persistence of this EntandoApp, allowed values: [mysql|oracle|postgresql|embedded|none]")
}
