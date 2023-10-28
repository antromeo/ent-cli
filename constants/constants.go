package constants

import "k8s.io/apimachinery/pkg/runtime/schema"

const (
	DefaultProfile          = "default"
	EntFolder               = ".ent"
	ProfilesFolder          = "profiles"
	GlobalConfigFileName    = "global-cfg"
	ConfigFile              = "cfg"
	ContainerRuntime        = "podman"
	EntandoBundleRepository = "docker.io/romeocontainers/entando-bundle-cli"
)

type Config struct {
	DesignedProfile string `yaml:"designedProfile"`
}

type ProfileConfig struct {
	AppName    string `yaml:"entandoAppName,omitempty""`
	Namespace  string `yaml:"entandoNamespace,omitempty""`
	DesignedVM string `yaml:"designedVM,omitempty"`
}

var EntandoDeBundleGroupVersionResource = schema.GroupVersionResource{
	Group:    "entando.org",
	Version:  "v1",
	Resource: "entandodebundles",
}
