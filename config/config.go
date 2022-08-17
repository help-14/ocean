package config

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/help-14/ocean-backup/services"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Services ConfigServices `yaml:"services"`
	Jobs     []BackupJob    `yaml:"jobs"`
}

type ConfigServices struct {
	Deta DetaConfig `yaml:"deta"`
}

type DetaConfig struct {
	ProjectKey string `yaml:"projectKey"`
	DrivePath  string `yaml:"drivePath"`
}

type BackupJob struct {
	Name     string               `yaml:"name"`
	UseZip   bool                 `yaml:"useZip"`
	Path     string               `yaml:"path"`
	UploadTo services.ServiceName `yaml:"uploadTo"`
}

func LoadConfig() Config {
	defaultConfig := Config{}

	yamlFile, err := ioutil.ReadFile(filepath.Join("data", "config.yaml"))
	if err != nil {
		fmt.Printf("Error reading config file: %s\n", err)
		return defaultConfig
	}

	var yamlConfig Config
	err = yaml.Unmarshal(yamlFile, &yamlConfig)
	if err != nil {
		fmt.Printf("Error parsing config file: %s\n", err)
		return defaultConfig
	}

	return yamlConfig
}
