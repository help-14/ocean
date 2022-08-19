package config

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Services ConfigServices `yaml:"services"`
	Jobs     []BackupJob    `yaml:"jobs"`
}

type ConfigServices struct {
	Deta       DetaConfig `yaml:"deta"`
	Cloudflare S3Config   `yaml:"cloudflare"`
}

func LoadConfig() (*Config, error) {
	yamlFile, err := ioutil.ReadFile(filepath.Join("config.yaml"))
	if err != nil {
		fmt.Printf("Error reading config file: %s\n", err)
		return nil, err
	}

	var yamlConfig *Config
	err = yaml.Unmarshal(yamlFile, &yamlConfig)
	if err != nil {
		fmt.Printf("Error parsing config file: %s\n", err)
		return nil, err
	}

	return yamlConfig, nil
}
