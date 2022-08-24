package services

import "github.com/help-14/ocean/config"

type Service interface {
	Name() string
	Setup(config config.ServiceConfig) error

	Connect() error
	Disconnect() error

	Upload(localPath string, remotePath string) error
}
