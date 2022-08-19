package services

import "github.com/help-14/ocean-backup/config"

type Service interface {
	Setup(config config.Config) error

	Connect() error
	Disconnect() error

	Upload(localPath string, remotePath string) error
}
