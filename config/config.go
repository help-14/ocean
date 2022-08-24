package config

import (
	"io/ioutil"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type Config struct {
	UseDashboard bool            `yaml:"useDashboard"`
	Services     []ServiceConfig `yaml:"services"`
	Jobs         []BackupJob     `yaml:"jobs"`
}

type ServiceConfig struct {
	Deta    DetaConfig `yaml:"deta"`
	S3      S3Config   `yaml:"s3"`
	Name    string     `yaml:"name"`
	Service string     `yaml:"service"`
}

func LoadConfig() (*Config, error) {
	yamlFile, err := ioutil.ReadFile(filepath.Join("config.yaml"))
	if err != nil {
		return nil, err
	}

	var yamlConfig *Config
	err = yaml.Unmarshal(yamlFile, &yamlConfig)
	if err != nil {
		return nil, err
	}

	return yamlConfig, nil
}
