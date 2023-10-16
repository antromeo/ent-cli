package cmd

import (
	. "ent-cli/constants"
	"ent-cli/profile"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
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

	viper.SetConfigName(filepath.Join(EntFolder, GlobalConfigFile))

	if err := viper.ReadInConfig(); err == nil {
		//fmt.Fprintln(os.Stderr, "Using global config file:", viper.ConfigFileUsed())
	}

	viper.SetConfigName(filepath.Join(EntFolder, ProfilesFolder, viper.GetString("designedProfile"), ConfigFile))

	viper.MergeInConfig()

	viper.AutomaticEnv()

}

func createEntDirectories() {
	// Perform directory creation or other installation tasks here.
	// For example, create a directory when the application is installed.
	home, _ := os.UserHomeDir()
	directoryPath := filepath.Join(home, EntFolder)

	// Check if the directory already exists.
	if _, err := os.Stat(directoryPath); os.IsNotExist(err) {
		// Directory doesn't exist, create it.
		profiles := filepath.Join(directoryPath, ProfilesFolder, DefaultProfile)
		err := os.MkdirAll(profiles, os.ModePerm)
		if err != nil {
			fmt.Printf("Error creating directory: %v\n", err)
			return
			//os.Exit(1)
		}

		profileConfig := ProfileConfig{
			AppName:    "quickstart",
			Namespace:  "entando",
			DesignedVM: "entando",
		}

		// Marshal the configuration to YAML format.
		yamlData, err := yaml.Marshal(&profileConfig)
		if err != nil {
			log.Fatalf("Error marshaling YAML: %v", err)
		}
		err = os.WriteFile(filepath.Join(profiles, ConfigFile+".yaml"), yamlData, 0644)
		if err != nil {
			log.Fatalf("Error writing to file: %v", err)
		}

		globalCfgFile := filepath.Join(directoryPath, GlobalConfigFile+".yaml")
		if _, err := os.Stat(globalCfgFile); os.IsNotExist(err) {
			_, err := os.Create(globalCfgFile)
			if err != nil {
				fmt.Printf("Error creating global-cfg file: %v\n", err)
				return
				//os.Exit(1)
			}

			globalConfig := Config{
				DesignedProfile: DefaultProfile,
			}

			// Marshal the configuration to YAML format.
			yamlData, err := yaml.Marshal(&globalConfig)
			if err != nil {
				log.Fatalf("Error marshaling YAML: %v", err)
			}
			err = os.WriteFile(globalCfgFile, yamlData, 0644)
			if err != nil {
				log.Fatalf("Error writing to file: %v", err)
			}

		} else if err != nil {
			// An error occurred while checking the directory existence.
			fmt.Printf("Error checking directory: %v\n", err)
			return
			//os.Exit(1)
		}
	}
}

type Config struct {
	DesignedProfile string `yaml:"designedProfile"`
}

func addSubCommands() {
	rootCmd.AddCommand(profile.ProfileCmd)
}
