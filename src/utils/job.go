package utils

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/help-14/ocean/config"
	"github.com/help-14/ocean/services"
	"github.com/robfig/cron/v3"
)

type JobRunner struct {
	Job     config.BackupJob
	Cron    *cron.Cron
	Service services.Service
}

func (runner *JobRunner) Run() error {
	uploadFile := ""
	uploadPath := runner.Job.Path

	stats, err := os.Stat(uploadPath)
	if os.IsNotExist(err) {
		return err
	}
	if !stats.IsDir() {
		uploadFile = stats.Name()
	}

	if (runner.Job.UseZip && !strings.Contains(uploadPath, ".zip")) || stats.IsDir() {
		uploadFile = createZipName(runner.Job.Name)
		uploadPath = createTempPath(uploadFile)
		err = CompressZip(runner.Job.Path, uploadPath)
		if err != nil {
			fmt.Println(err.Error())
			return err
		}
	}

	if runner.Service == nil {
		return errors.New("No service found with this name: " + runner.Job.UploadTo)
	}
	remotePath := createRemotePath(uploadFile)
	err = runner.Service.Upload(uploadPath, remotePath)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	if strings.Contains(uploadPath, tempFolder()) {
		return os.Remove(uploadPath)
	}
	return nil
}

func createZipName(name string) string {
	result := strings.ReplaceAll(name, " ", "_")
	result = strings.ToLower(result)
	result = strings.TrimSpace(result)
	result = result + ".zip"
	return result
}

func tempFolder() string {
	tempPath := os.Getenv("TEMP_PATH")

	if len(tempPath) <= 0 {
		pwd, _ := os.Getwd()
		tempPath = filepath.Join(pwd, "temp")
	}

	if _, err := os.Stat(tempPath); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(tempPath, os.ModePerm)
		if err != nil {
			log.Println(err)
		}
	}
	return tempPath
}

func createTempPath(fileName string) string {
	return filepath.Join(tempFolder(), fileName)
}

func createRemotePath(fileName string) string {
	return filepath.Join(time.Now().Format("2006-01-02"), fileName)
}
