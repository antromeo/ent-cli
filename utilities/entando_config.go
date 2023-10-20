package utilities

import (
	. "ent-cli/constants"
	"fmt"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

type EntandoConfig struct {
	entandoHomeDir string
	extension      string
}

var instance *EntandoConfig
var once sync.Once

func GetEntandoConfigInstance() *EntandoConfig {
	once.Do(func() {
		home, _ := os.UserHomeDir()
		instance = &EntandoConfig{
			entandoHomeDir: home,
			extension:      "yaml",
		}
	})
	return instance
}

func (e *EntandoConfig) SetHomeDir(homeDir string) {
	e.entandoHomeDir = homeDir
}

func (e *EntandoConfig) GetHomeDir() string {
	return e.entandoHomeDir
}

func (ec *EntandoConfig) GetEntFolderFilePath() string {
	return filepath.Join(ec.entandoHomeDir, EntFolder)
}

func (ec *EntandoConfig) GetEntGlobalConfigFilePath() string {
	globalCfg := strings.Join([]string{GlobalConfigFileName, ec.extension}, ".")
	return filepath.Join(ec.GetEntFolderFilePath(), globalCfg)
}

func (ec *EntandoConfig) GetProfilesFilePath() string {
	return filepath.Join(ec.GetEntFolderFilePath(), ProfilesFolder)
}

func (ec *EntandoConfig) GetProfileFilePath(profile string) string {
	return filepath.Join(ec.GetProfilesFilePath(), profile)
}

func (ec *EntandoConfig) GetEntConfigFilePathByProfile(profile string) string {
	cfg := strings.Join([]string{ConfigFile, ec.extension}, ".")
	return filepath.Join(ec.GetProfileFilePath(profile), cfg)
}

func (ec *EntandoConfig) GetProfile() string {
	return viper.GetString("designedProfile")
}

func (ec *EntandoConfig) GetProfilesDirectories() ([]os.DirEntry, error) {
	files, err := os.ReadDir(ec.GetProfilesFilePath())
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}
	return files, nil
}

func (ec *EntandoConfig) GetProfiles() ([]string, error) {
	files, err := os.ReadDir(ec.GetProfilesFilePath())
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}
	var dirNames []string

	// Iterate over the DirEntry objects and extract directory names
	for _, entry := range files {
		if entry.IsDir() {
			dirNames = append(dirNames, entry.Name())
		}
	}
	return dirNames, nil
}
