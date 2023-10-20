package utilities

import (
	"path/filepath"
	"testing"
)

var customDir string
var entandoConfig *EntandoConfig

func init() {
	customDir = "/home/user"
	entandoConfig = GetEntandoConfigInstance()
	entandoConfig.SetHomeDir(customDir)
}

func TestGetEntFolderFilePath(t *testing.T) {
	filePath := entandoConfig.GetEntFolderFilePath()

	expectedFilePath := filepath.Join(customDir, ".ent")

	if filePath != expectedFilePath {
		t.Errorf("Expected file path %s, but got %s", expectedFilePath, filePath)
	}
}

func TestGetEntGlobalConfigFilePath(t *testing.T) {
	filePath := entandoConfig.GetEntGlobalConfigFilePath()

	expectedFilePath := filepath.Join(customDir, ".ent", "global-cfg.yaml")

	if filePath != expectedFilePath {
		t.Errorf("Expected file path %s, but got %s", expectedFilePath, filePath)
	}
}

func TestGetProfilesFilePath(t *testing.T) {
	filePath := entandoConfig.GetProfilesFilePath()

	expectedFilePath := filepath.Join(customDir, ".ent", "profiles")

	if filePath != expectedFilePath {
		t.Errorf("Expected file path %s, but got %s", expectedFilePath, filePath)
	}
}

func TestGetProfileFilePath2(t *testing.T) {
	filePath := entandoConfig.GetProfileFilePath("bob")
	expectedFilePath := filepath.Join(customDir, ".ent", "profiles", "bob")

	if filePath != expectedFilePath {
		t.Errorf("Expected file path %s, but got %s", expectedFilePath, filePath)
	}
}

func TestGetEntConfigFilePathByProfile2(t *testing.T) {
	filePath := entandoConfig.GetEntConfigFilePathByProfile("bob")
	expectedFilePath := filepath.Join(customDir, ".ent", "profiles", "bob", "cfg.yaml")

	if filePath != expectedFilePath {
		t.Errorf("Expected file path %s, but got %s", expectedFilePath, filePath)
	}
}
