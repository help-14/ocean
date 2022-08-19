package services

import (
	"bufio"
	"os"

	"github.com/deta/deta-go/deta"
	"github.com/deta/deta-go/service/drive"
	obconfig "github.com/help-14/ocean-backup/config"
)

type DetaService struct {
	config obconfig.DetaConfig
	drive  *drive.Drive
}

func (service DetaService) Setup(config obconfig.Config) error {
	service.config = config.Services.Deta
	return nil
}

func (service DetaService) Connect() error {
	d, err := deta.New(deta.WithProjectKey(service.config.ProjectKey))
	if err != nil {
		return err
	}

	drive, err := drive.New(d, service.config.DriveName)
	if err != nil {
		return err
	}

	service.drive = drive
	return nil
}

func (service DetaService) Disconnect() error {
	service.drive = nil
	return nil
}

func (service DetaService) Upload(localPath string, remotePath string) error {
	file, err := os.Open(localPath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = service.drive.Put(&drive.PutInput{
		Name: remotePath,
		Body: bufio.NewReader(file),
	})
	if err != nil {
		return err
	}

	return nil
}
