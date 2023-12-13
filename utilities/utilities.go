package utilities

import (
	"bufio"
	"context"
	"fmt"
	"github.com/antromeo/ent-cli/v2/types"
	"github.com/spf13/viper"
	"io"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"net/http"
	"os"
	"strings"
)

func ReadString(inputText string, required bool) string {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print(inputText + ": ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	if required {
		for len(input) == 0 {
			fmt.Print("Input cannot be empty. Please enter a non-empty string: ")
			input, _ = reader.ReadString('\n')
			input = strings.TrimSpace(input)
		}
	}

	return input
}

func IsEmpty(str string) bool {
	if len(str) == 0 {
		return true
	}
	return false
}

func ShowAdditionalCommandsInHelp() {
	fmt.Println("ADDITIONAL COMMANDS")
	fmt.Println("  deploy       Generates the CR and deploys it to the currently attached EntandoApp")
	fmt.Println("  install      Installs into currently attached EntandoApp the bundle in the current directory")
}

func HttpGet(url string) []byte {
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error making HTTP request: ", err)
		os.Exit(1)
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		fmt.Println("HTTP request failed with status code: ", response.Status)
		os.Exit(1)
	}
	data, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading HTTP response: ", err)
		os.Exit(1)
	}
	return data
}

func RetrieveProtocol() string {
	k8sClient := GetKubeClientInstance()
	var forceProtocol = viper.GetString("FORCE_URL_SCHEME")
	if len(forceProtocol) > 0 && (forceProtocol == "http" || forceProtocol == "https") {
		return forceProtocol
	}

	appName := viper.GetString("EntandoAppName")
	ingressName := strings.Join([]string{appName, "ingress"}, "-")
	ingress, _ := k8sClient.ClientSet.NetworkingV1().Ingresses(k8sClient.Namespace).Get(context.TODO(), ingressName, metav1.GetOptions{})
	if len(ingress.Spec.TLS) > 0 {
		return "https"
	} else {
		return "http"
	}
}

func GetEntandoIngressesPath(protocol string) *types.EntandoIngressesPath {
	k8sClient := GetKubeClientInstance()
	appName := viper.GetString("EntandoAppName")
	ingressName := strings.Join([]string{appName, "ingress"}, "-")
	ingressResource, err := k8sClient.ClientSet.NetworkingV1().Ingresses(k8sClient.Namespace).Get(context.TODO(), ingressName, metav1.GetOptions{})
	if err != nil {
		fmt.Printf("Error checking ingress: %v\n", err)
		os.Exit(1)
	}
	hostname := protocol + "://" + ingressResource.Spec.Rules[0].Host

	return &types.EntandoIngressesPath{
		AppBuilder: strings.Join([]string{hostname, "app-builder"}, "/"),
		Ecr:        strings.Join([]string{hostname, "digital-exchange"}, "/"),
	}
}
