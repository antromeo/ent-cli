package utilities

import (
	. "ent-cli/constants"
	"fmt"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

func GetProfile() string {
	return viper.GetString("designedProfile")
}
func GetProfiles() ([]os.DirEntry, error) {
	files, err := os.ReadDir(GetProfilesFilePath())
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}
	return files, nil
}

func GetEntFilePath() string {
	home, _ := os.UserHomeDir()
	return filepath.Join(home, EntFolder)
}

func GetProfilesFilePath() string {
	home, _ := os.UserHomeDir()
	return filepath.Join(home, EntFolder, ProfilesFolder)
}

func GetProfileFilePath(profile string) string {
	home, _ := os.UserHomeDir()
	return filepath.Join(home, EntFolder, ProfilesFolder, profile)
}
