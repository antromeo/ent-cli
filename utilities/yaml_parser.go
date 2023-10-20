package utilities

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

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

func ReadFileToYaml(filePath string, data interface{}) {
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Error reading the file: %v", err)
	}
	if err := yaml.Unmarshal(fileContent, data); err != nil {
		log.Fatalf("Error unmarshaling YAML: %v", err)
	}
}
