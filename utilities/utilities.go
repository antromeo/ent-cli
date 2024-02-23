package utilities

import (
	"bufio"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/antromeo/ent-cli/v2/types"
	"github.com/spf13/viper"
	"io"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"net/http"
	"os"
	"regexp"
	"runtime"
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

func IsValidURL(input string) bool {
	regexPattern := `^[^:]*://[-A-Za-z0-9\+&@#/%?=~_|!:,.;]*[-A-Za-z0-9\+&@#/%=~_|]`
	match, err := regexp.MatchString(regexPattern, input)
	if err != nil {
		fmt.Println("Error compiling regex:", err)
		return false
	}

	return match
}

func HashAndTruncate(input string) string {
	hash := sha256.New()
	hash.Write([]byte(input))
	hashString := hex.EncodeToString(hash.Sum(nil))
	return hashString[:8]
}

func NormalizeName(input string) string {
	pattern := "[_:./]"
	regex := regexp.MustCompile(pattern)
	return regex.ReplaceAllString(input, "-")
}

func getOS() string {
	switch os := runtime.GOOS; os {
	case "linux":
		return "linux"
	case "windows":
		return "win"
	case "darwin":
		return "macos"
	default:
		return "unknown"
	}
}

func getArchitecture() string {
	switch arch := runtime.GOARCH; arch {
	case "amd64":
		return "x64"
	case "arm64":
		return "arm64"
	default:
		return "unknown"
	}
}

// TODO: make multiplatform and multiarch
type SourceBinary struct {
	Name     string
	Path     string
	Tag      string
	Sha      string
	RepoName string
}

func (sb *SourceBinary) DetermineUrl() string {
	return fmt.Sprintf("https://github.com/antromeo/%s/releases/download/%s/entando-bundle-cli-node14-%s-%s-%s", sb.RepoName, sb.Tag, getOS(), getArchitecture(), sb.Sha)
}

var EntBundleCliBinary = SourceBinary{
	Name:     "entando-bundle-cli",
	Path:     GetEntandoConfigInstance().GetEntBundleCliBinFilePath(),
	Tag:      "v1.1.2",
	Sha:      "2eb3bb76d2c86fc825957893620c2e05d3b20f5c",
	RepoName: "entando-bundle-cli",
}

var EntExtBinaries = []SourceBinary{EntBundleCliBinary}
