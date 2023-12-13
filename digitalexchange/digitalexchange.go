package digitalexchange

import (
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

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Printf("\nerror creating request: %v\n", err)
		os.Exit(1)
	}
	bearer := "Bearer " + token
	req.Header.Set("Authorization", bearer)

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		fmt.Printf("\nerror making request: %v\n", err)
		os.Exit(1)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("\nerror parsing response: %v\n", err)
		os.Exit(1)
	}

	var ecrComponentsResponse types.EcrComponentsResponse

	err = json.Unmarshal(body, &ecrComponentsResponse)

	return ecrComponentsResponse

}
