package types

type TokenEndpointPayload struct {
	TokenEndpoint string `json:"token_endpoint"`
}

type EntandoIngressesPath struct {
	AppBuilder string
	Ecr        string
}

type ClientSecret struct {
	ClientId     string `json:"clientId"`
	ClientSecret string `json:"clientSecret"`
	AuthUrl      string `json:"authUrl"`
	Realm        string `json:"realm"`
}

type EcrComponentsResponse struct {
	Payload []EcrComponentsPayload `json:"payload"`
}

type EcrComponentsPayload struct {
	Code           string   `json:"code"`
	Title          string   `json:"title"`
	Description    string   `json:"description"`
	RepoUrl        string   `json:"repoUrl"`
	BundleType     string   `json:"bundleType"`
	Thumbnail      string   `json:"thumbnail"`
	ComponentTypes []string `json:"componentTypes"`
	LastJob        LastJob  `json:"lastJob"`
}

type LastJob struct {
	ID               string     `json:"id"`
	ComponentID      string     `json:"componentId"`
	ComponentName    string     `json:"componentName"`
	ComponentVersion string     `json:"componentVersion"`
	Progress         float64    `json:"progress"`
	Status           string     `json:"status"`
	ComponentJobs    []struct{} `json:"componentJobs"`
}

type EcrComponentUninstallResponse struct {
	Payload EcrComponentUninstall `json:"payload"`
}

type EcrComponentUninstall struct {
	ID               string  `json:"id"`
	ComponentID      string  `json:"componentId"`
	ComponentName    string  `json:"componentName"`
	ComponentVersion string  `json:"componentVersion"`
	Progress         float64 `json:"progress"`
	Status           string  `json:"status"`
}
