package cmd

import (
	"bytes"
	"fmt"
	"github.com/antromeo/ent-cli/v2/cmd/app"
	"github.com/antromeo/ent-cli/v2/cmd/config"
	"github.com/antromeo/ent-cli/v2/cmd/ecr"
	"github.com/antromeo/ent-cli/v2/cmd/profile"
	"github.com/antromeo/ent-cli/v2/cmd/quickstart"
	"github.com/antromeo/ent-cli/v2/cmd/tenant"
	. "github.com/antromeo/ent-cli/v2/constants"
	"github.com/antromeo/ent-cli/v2/utilities"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:     "ent-cli",
	Short:   "A new version of ent",
	Long:    `A new version of ent`,
	Version: "2.0.0",
}

func Execute() {
	args := os.Args[1:]
	if _, _, err := rootCmd.Find(args); err != nil && len(args) > 0 {
		// not found, search executable in plugins
		if pathPlugin, err := findEntPlugin(args); err != nil {
			fmt.Printf("Error: unknown command \"%v\" for \"%v\"\n", args[0], rootCmd.Name())
			os.Exit(1)
		} else {
			// run plugin
			runEntPlugin(pathPlugin, args)
		}

	} else {
		err := rootCmd.Execute()
		if err != nil {
			os.Exit(1)
		}
	}

}

func init() {
	createEntDirectories()
	downloadBinaries()
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file")

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	addSubCommands()
}

func initConfig() {
	entandoConfig := utilities.GetEntandoConfigInstance()

	viper.AddConfigPath(entandoConfig.GetHomeDir())

	viper.SetConfigName(filepath.Join(EntFolder, GlobalConfigFileName))
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	viper.SetConfigName(filepath.Join(EntFolder, ProfilesFolder, entandoConfig.GetProfile(), ConfigFile))
	if err := viper.MergeInConfig(); err != nil {
		log.Fatalf("Error reading profile config file: %v", err)
	}

	viper.AutomaticEnv()

}

func createEntDirectories() {
	entandoConfig := utilities.GetEntandoConfigInstance()
	entFolderFilePath := entandoConfig.GetEntFolderFilePath()

	// Check if the directory already exists.
	if _, err := os.Stat(entFolderFilePath); os.IsNotExist(err) {
		// Directory doesn't exist, create the default profile files.
		createDefaultProfileDirectories()
		createGlobalConfigFile()
	}
	if _, err := os.Stat(entandoConfig.GetEntBinFolderFilePath()); os.IsNotExist(err) {
		createBinFolder()
	}

}

func createDefaultProfileDirectories() {
	entandoConfig := utilities.GetEntandoConfigInstance()
	defaultProfileFilePath := entandoConfig.GetProfileFilePath(DefaultProfile)

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

	utilities.WriteYamlToFile(entandoConfig.GetEntConfigFilePathByProfile(DefaultProfile), defaultProfileConfig)
}

func createBinFolder() {
	entandoConfig := utilities.GetEntandoConfigInstance()
	entBinFilePath := entandoConfig.GetEntBinFolderFilePath()

	err := os.Mkdir(entBinFilePath, 0770)
	if err != nil {
		fmt.Printf("Error creating directory: %v\n", err)
		return
	}
}
func createGlobalConfigFile() {
	entandoConfig := utilities.GetEntandoConfigInstance()
	globalCfgFilePath := entandoConfig.GetEntGlobalConfigFilePath()

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
	rootCmd.AddCommand(quickstart.QuickstartCmd)
	rootCmd.AddCommand(app.AppCmd)
	rootCmd.AddCommand(tenant.TenantCmd)
	rootCmd.AddCommand(ecr.EcrCmd)
	rootCmd.AddCommand(config.ConfigCmd)
}

func findEntPlugin(args []string) (string, error) {
	pluginName := strings.Join([]string{rootCmd.Name(), args[0]}, "-")
	if pathPlugin, err := exec.LookPath(pluginName); err != nil {
		return "", err
	} else {
		return pathPlugin, nil
	}
}

func runEntPlugin(pathPlugin string, args []string) {
	cmd := exec.Command(pathPlugin, args[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func downloadBinaries() {
	//entandoConfig := utilities.GetEntandoConfigInstance()
	//entBinFilePath := entandoConfig.GetEntBinFolderFilePath()
	for _, extBin := range utilities.EntExtBinaries {
		// TODO: make for all binaries
		//entBundleCliPath := path.Join(entBinFilePath, "entando-bundle-cli")
		if _, err := os.Stat(extBin.Path); os.IsNotExist(err) {
			// TODO: url dynamic and print a message of waiting, change repository, make the same for entando-bundler (use node-17)
			// TODO: when kubectl is not aligned, the errors are not good
			// TODO: set var as version for bundle-cli
			// TODO: store the files in profile bin folder (to create)
			fmt.Println("Loading binaries...")
			url := extBin.DetermineUrl()
			response := utilities.HttpGet(url)

			// Create the output file
			out, err := os.Create(extBin.Path)
			out.Chmod(0777)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			defer out.Close()

			// Copy the response body to the output file
			_, err = io.Copy(out, bytes.NewReader(response))
			if err != nil {
				fmt.Println("Error:", err)
				return
			}

			fmt.Println("File downloaded successfully.")
		}
	}

}
