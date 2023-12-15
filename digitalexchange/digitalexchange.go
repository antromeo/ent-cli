package digitalexchange

import (
	"bytes"
	"fmt"
	"github.com/antromeo/ent-cli/v2/constants"
	"github.com/antromeo/ent-cli/v2/idp"
	"github.com/antromeo/ent-cli/v2/types"
	"github.com/antromeo/ent-cli/v2/utilities"
	"io"
	"k8s.io/apimachinery/pkg/util/json"
	"net/http"
	"os"
	"strings"
)

func GetEcrComponents() types.EcrComponentsResponse {
	clientParams := idp.RetrieveClientParameters()
	protocol := utilities.RetrieveProtocol()
	token := idp.GetToken(protocol, clientParams)
	ingressesPaths := utilities.GetEntandoIngressesPath(protocol)

	url := strings.Join([]string{ingressesPaths.Ecr, constants.EcrApis.Components}, "/")

	body := makeEcrRequest(http.MethodGet, url, token, nil)

	var ecrComponentsResponse types.EcrComponentsResponse

	err := json.Unmarshal(body, &ecrComponentsResponse)

	if err != nil {
		fmt.Printf("\nerror unmarshalling data: %v\n", err)
		os.Exit(1)
	}

	return ecrComponentsResponse

}

func InstallComponent(bundleCode string, version string, strategy string) types.EcrComponentUninstallResponse {
	clientParams := idp.RetrieveClientParameters()
	protocol := utilities.RetrieveProtocol()
	token := idp.GetToken(protocol, clientParams)
	ingressesPaths := utilities.GetEntandoIngressesPath(protocol)

	path := strings.Replace(constants.EcrApis.Install, "{component}", bundleCode, 1)
	endpoint := strings.Join([]string{ingressesPaths.Ecr, path}, "/")

	inputBody := map[string]interface{}{
		"version":          version,
		"conflictStrategy": strategy,
	}
	responseBody := makeEcrRequest(http.MethodPost, endpoint, token, inputBody)

	var ecrComponentUninstallResponse types.EcrComponentUninstallResponse

	err := json.Unmarshal(responseBody, &ecrComponentUninstallResponse)
	if err != nil {
		fmt.Printf("\nerror unmarshalling data: %v\n", err)
		os.Exit(1)
	}

	return ecrComponentUninstallResponse

}
func UninstallComponent(bundleCode string) types.EcrComponentUninstallResponse {
	clientParams := idp.RetrieveClientParameters()
	protocol := utilities.RetrieveProtocol()
	token := idp.GetToken(protocol, clientParams)
	ingressesPaths := utilities.GetEntandoIngressesPath(protocol)

	path := strings.Replace(constants.EcrApis.Uninstall, "{component}", bundleCode, 1)
	url := strings.Join([]string{ingressesPaths.Ecr, path}, "/")

	body := makeEcrRequest(http.MethodPost, url, token, nil)

	var ecrComponentUninstallResponse types.EcrComponentUninstallResponse

	err := json.Unmarshal(body, &ecrComponentUninstallResponse)
	if err != nil {
		fmt.Printf("\nerror unmarshalling data: %v\n", err)
		os.Exit(1)
	}

	return ecrComponentUninstallResponse

}

func makeEcrRequest(method string, url string, token string, inputbody map[string]interface{}) []byte {
	payloadBytes, err := json.Marshal(inputbody)
	if err != nil {
		fmt.Printf("Error encoding JSON: %v\n", err)
		os.Exit(1)
	}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(payloadBytes))
	if err != nil {
		fmt.Printf("\nerror creating request: %v\n", err)
		os.Exit(1)
	}

	bearer := "Bearer " + token
	req.Header.Set("Authorization", bearer)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("\nerror making request: %v\n", err)
		os.Exit(1)
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusAccepted {
		fmt.Println("HTTP request failed with status code: ", resp.StatusCode)
		os.Exit(1)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("\nerror parsing response: %v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()
	return body
}
