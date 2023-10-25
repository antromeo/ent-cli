package constants

const (
	DefaultProfile       = "default"
	EntFolder            = ".ent"
	ProfilesFolder       = "profiles"
	GlobalConfigFileName = "global-cfg"
	ConfigFile           = "cfg"
	ContainerRuntime     = "podman"
)

type Config struct {
	DesignedProfile string `yaml:"designedProfile"`
}

type ProfileConfig struct {
	AppName    string `yaml:"entandoAppName,omitempty""`
	Namespace  string `yaml:"entandoNamespace,omitempty""`
	DesignedVM string `yaml:"designedVM,omitempty"`
}
