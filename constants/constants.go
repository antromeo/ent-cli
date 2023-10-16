package constants

const (
	DefaultProfile   = "default"
	EntFolder        = ".ent"
	ProfilesFolder   = "profiles"
	GlobalConfigFile = "global-cfg"
	ConfigFile       = "cfg"
)

type ProfileConfig struct {
	AppName    string `yaml:"entandoAppName,omitempty""`
	Namespace  string `yaml:"entandoNamespace,omitempty""`
	DesignedVM string `yaml:"designedVM,omitempty"`
}
