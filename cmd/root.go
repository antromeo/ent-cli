package cmd

import (
	. "ent-cli/constants"
	"ent-cli/profile"
	"ent-cli/utilities"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"os"
	"path/filepath"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:     "ent-cli",
	Short:   "A new version of ent",
	Long:    `A new version of ent`,
	Version: "2.0.0",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	createEntDirectories()
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file")

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	addSubCommands()
}

func initConfig() {
	home, err := os.UserHomeDir()
	cobra.CheckErr(err)

	viper.AddConfigPath(home)
	viper.SetConfigType("yaml")

	viper.SetConfigName(filepath.Join(EntFolder, GlobalConfigFileName))

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	viper.SetConfigName(filepath.Join(EntFolder, ProfilesFolder, viper.GetString("designedProfile"), ConfigFile))

	viper.MergeInConfig()

	viper.AutomaticEnv()

}

func createEntDirectories() {

	entFolderFilePath := utilities.GetEntFolderFilePath()

	// Check if the directory already exists.
	if _, err := os.Stat(entFolderFilePath); os.IsNotExist(err) {
		// Directory doesn't exist, create the default profile files.
		createDefaultProfileDirectories()
		createGlobalConfigFile()

	}
}

func createDefaultProfileDirectories() {
	defaultProfileFilePath := utilities.GetProfileFilePath(DefaultProfile)

	err := os.MkdirAll(defaultProfileFilePath, 0770)
	if err != nil {
		fmt.Printf("Error creating directory: %v\n", err)
		return
	}

	defaultProfileConfig := ProfileConfig{
		AppName:    "quickstart",
		Namespace:  "entando",
		DesignedVM: "entando",
	}

	utilities.WriteYamlToFile(utilities.GetEntConfigFilePathByProfile(DefaultProfile), defaultProfileConfig)
}

func createGlobalConfigFile() {
	globalCfgFilePath := utilities.GetEntGlobalConfigFilePath()

	if _, err := os.Stat(globalCfgFilePath); os.IsNotExist(err) {
		globalConfig := Config{
			DesignedProfile: DefaultProfile,
		}

		utilities.WriteYamlToFile(globalCfgFilePath, globalConfig)

	} else if err != nil {
		fmt.Printf("Error checking directory: %v\n", err)
		os.Exit(1)
	}
}

func addSubCommands() {
	rootCmd.AddCommand(profile.ProfileCmd)
}
