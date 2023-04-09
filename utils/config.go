package utils

import (
	"io"
	"os"

	"github.com/anyingiit/GoReactResourceManagement/structs"
	"gopkg.in/yaml.v2"
)

func ReadConfigFile(configFilePath string) (*structs.ProjectConfig, error) {
	file, err := os.Open(configFilePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	yamlFile, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	// Parse YAML file
	var config structs.ProjectConfig
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
