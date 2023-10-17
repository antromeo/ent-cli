package utilities

import (
	. "ent-cli/constants"
	"fmt"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func GetProfile() string {
	return viper.GetString("designedProfile")
}

func GetProfilesDirectories() ([]os.DirEntry, error) {
	files, err := os.ReadDir(GetProfilesFilePath())
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}
	return files, nil
}

func GetProfiles() ([]string, error) {
	files, err := os.ReadDir(GetProfilesFilePath())
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

func GetEntFolderFilePath() string {
	home, _ := os.UserHomeDir()
	return filepath.Join(home, EntFolder)
}

func GetEntConfigFilePathByProfile(profile string) string {
	extension := "yaml"
	cfg := strings.Join([]string{ConfigFile, extension}, ".")
	return filepath.Join(GetProfileFilePath(profile), cfg)
}

func GetEntGlobalConfigFilePath() string {
	home, _ := os.UserHomeDir()
	extension := "yaml"
	globalCfg := strings.Join([]string{GlobalConfigFileName, extension}, ".")
	return filepath.Join(home, EntFolder, globalCfg)
}

func GetProfilesFilePath() string {
	home, _ := os.UserHomeDir()
	return filepath.Join(home, EntFolder, ProfilesFolder)
}

func GetProfileFilePath(profile string) string {
	home, _ := os.UserHomeDir()
	return filepath.Join(home, EntFolder, ProfilesFolder, profile)
}

func WriteYamlToFile(filePath string, data interface{}) {
	yamlData, err := yaml.Marshal(&data)
	if err != nil {
		log.Fatalf("Error marshaling YAML: %v", err)
	}
	err = os.WriteFile(filePath, yamlData, 0600)
	if err != nil {
		log.Fatalf("Error writing to file: %v", err)
	}
}
