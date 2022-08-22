package utils

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/help-14/ocean-backup/config"
	"github.com/help-14/ocean-backup/services"
	"github.com/robfig/cron/v3"
)

type JobRunner struct {
	Job     config.BackupJob
	Cron    *cron.Cron
	Service services.Service
}

func (runner JobRunner) Run() error {
	uploadFile := ""
	uploadPath := runner.Job.Path

	stats, err := os.Stat(uploadPath)
	if os.IsNotExist(err) {
		return err
	}

	if runner.Job.UseZip || stats.IsDir() {
		pwd, _ := os.Getwd()
		uploadFile = createZipName(runner.Job.Name)
		uploadPath = filepath.Join(pwd, "temp", uploadFile)
		err = ZipFolder(runner.Job.Path, uploadFile)
		if err != nil {
			return err
		}
	}

	if runner.Service == nil {
		return errors.New("No service found with this name: " + runner.Job.UploadTo)
	}
	remotePath := createRemotePath(uploadFile)
	return runner.Service.Upload(uploadPath, remotePath)
}

func createZipName(name string) string {
	result := strings.ReplaceAll(name, " ", "_")
	result = strings.ToLower(result)
	result = strings.TrimSpace(result)
	if !strings.Contains(result, ".zip") {
		result = result + ".zip"
	}
	return result
}

func createRemotePath(fileName string) string {
	return filepath.Join("/", time.Now().Format("2006-02-13"), fileName)
}
