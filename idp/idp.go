package idp

import (
	"context"
	"fmt"
	"github.com/antromeo/ent-cli/v2/types"
	"github.com/antromeo/ent-cli/v2/utilities"
	"github.com/spf13/viper"
	"io"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/json"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func RetrieveClientParameters() types.ClientSecret {
	k8sClient := utilities.GetKubeClientInstance()

	// external
	externalKeycloakSecret, err := k8sClient.ClientSet.CoreV1().Secrets(k8sClient.Namespace).Get(context.TODO(), "external-sso-secret", metav1.GetOptions{})
	if err == nil {
		return types.ClientSecret{
			ClientId:     string(externalKeycloakSecret.Data["clientId"]),
			ClientSecret: string(externalKeycloakSecret.Data["clientSecret"]),
			AuthUrl:      string(externalKeycloakSecret.Data["authUrl"]),
			Realm:        string(externalKeycloakSecret.Data["realm"]),
		}
	}

	// internal
	appName := viper.GetString("EntandoAppName")
	deSecretName := strings.Join([]string{appName, "de-secret"}, "-")
	keycloakSecret, _ := k8sClient.ClientSet.CoreV1().Secrets(k8sClient.Namespace).Get(context.TODO(), deSecretName, metav1.GetOptions{})

	ssoIngress, _ := k8sClient.ClientSet.NetworkingV1().Ingresses(k8sClient.Namespace).Get(context.TODO(), "default-sso-in-namespace-ingress", metav1.GetOptions{})

	return types.ClientSecret{
		ClientId:     string(keycloakSecret.Data["clientId"]),
		ClientSecret: string(keycloakSecret.Data["clientSecret"]),
		AuthUrl:      strings.Join([]string{ssoIngress.Spec.Rules[0].Host, "auth"}, "/"),
		Realm:        "entando",
	}

}

func GetTokenEndpoint(protocol string, authUrl string, realm string) types.TokenEndpointPayload {
	authUrl = strings.Join([]string{protocol, authUrl}, "://")
	url := strings.Join([]string{authUrl, "realms", realm, ".well-known/openid-configuration"}, "/")
	res, err := http.Get(url)
	if err != nil {
		fmt.Printf("\nerror making http request: %s\n", err)
		os.Exit(1)
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		fmt.Printf("\nerror making http request, status code: %v\n", res.StatusCode)
		os.Exit(1)
	}

	var tokenEndpoint types.TokenEndpointPayload

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("\nError parsing response: %v\n", err)
		os.Exit(1)
	}

	err = json.Unmarshal(resBody, &tokenEndpoint)
	if err != nil {
		fmt.Printf("\nError decoding JSON: %v\n", err)
		os.Exit(1)
	}

	return tokenEndpoint
}

func GetToken(protocol string, clientParams types.ClientSecret) string {

	tokenPayload := GetTokenEndpoint(protocol, clientParams.AuthUrl, clientParams.Realm)

	bodyReq := url.Values{
		"grant_type": {"client_credentials"},
	}

	request, err := http.NewRequest(http.MethodPost, tokenPayload.TokenEndpoint, strings.NewReader(bodyReq.Encode()))
	request.SetBasicAuth(clientParams.ClientId, clientParams.ClientSecret)
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	res, err := client.Do(request)
	if err != nil {
		fmt.Printf("\nerror making http request: %s\n", err)
		os.Exit(1)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		fmt.Printf("\nerror making http request, status code: %v\n", res.StatusCode)
		os.Exit(1)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Errorf("\nerror parsing response: %v\n", err)
		os.Exit(1)
	}
	var responseMap map[string]interface{}
	err = json.Unmarshal(body, &responseMap)
	if err != nil {
		fmt.Errorf("\nerror unmarshalling response: %v\n", err)
		os.Exit(1)
	}
	return responseMap["access_token"].(string)
}
