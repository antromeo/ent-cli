package constants

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
)

const (
	DefaultProfile       = "default"
	EntFolder            = ".ent"
	ProfilesFolder       = "profiles"
	GlobalConfigFileName = "global-cfg"
	ConfigFile           = "cfg"
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

var EntandoResourcesTemplates = []string{
	"https://raw.githubusercontent.com/entando/entando-releases/%s/dist/ge-1-1-6/namespace-scoped-deployment/cluster-resources.yaml",
	"https://raw.githubusercontent.com/entando/entando-releases/%s/dist/ge-1-1-6/namespace-scoped-deployment/namespace-resources.yaml",
	"https://raw.githubusercontent.com/entando/entando-releases/%s/dist/ge-1-1-6/samples/entando-operator-config.yaml",
}

type EcrApisType struct {
	Components string
	Install    string
	Uninstall  string
}

var EcrApis = EcrApisType{
	Components: "components",
	Install:    "components/{component}/install",
	Uninstall:  "components/{component}/uninstall",
}
