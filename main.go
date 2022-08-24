package main

import (
	"bufio"
	"log"
	"os"
	"strings"

	"github.com/help-14/ocean-backup/config"
	"github.com/help-14/ocean-backup/dashboard"
	"github.com/help-14/ocean-backup/services"
	"github.com/help-14/ocean-backup/utils"
	"github.com/robfig/cron/v3"
)

var LoadedConfig config.Config
var SupportServices []services.Service
var JobRunners []utils.JobRunner

func main() {
	loadedConfig, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Reading config.yaml failed!\n%s\n", err.Error())
	}
	LoadedConfig = *loadedConfig

	for _, serviceConfig := range LoadedConfig.Services {
		var newService services.Service

		switch serviceConfig.Service {
		case config.ServiceName_Deta:
			newService = new(services.DetaService)
		case config.ServiceName_Cloudflare:
			newService = new(services.CloudFlareService)
		default:
			log.Fatalf("Unsupported service name: %s\n%s\n", serviceConfig.Service, err.Error())
		}

		err = newService.Setup(serviceConfig)
		if err != nil {
			log.Fatalf("Setup service '%s' failed!\n%s\n", serviceConfig.Name, err.Error())
		}
		err = newService.Connect()
		if err != nil {
			log.Fatalf("Connect to service '%s' failed!\n%s\n", serviceConfig.Name, err.Error())
		}
		log.Printf("Service '%s' is connected!\n", serviceConfig.Name)
		SupportServices = append(SupportServices, newService)
	}

	for _, job := range LoadedConfig.Jobs {
		var runner utils.JobRunner
		runner.Job = job

		cronTime := job.RunAt
		if len(strings.TrimSpace(cronTime)) == 0 {
			cronTime = "0 0 * * *"
		}
		cron := cron.New()
		cron.AddFunc(cronTime, func() {
			err = runner.Run()
			if err != nil {
				log.Println(err.Error())
			} else {
				log.Printf("'%s' uploaded to '%s'\n", runner.Job.Path, runner.Service.Name())
			}
		})
		cron.Start()
		runner.Cron = cron

		for _, service := range SupportServices {
			if service.Name() == job.UploadTo {
				runner.Service = service
			}
		}

		JobRunners = append(JobRunners, runner)
	}

	log.Println("Ocean backup is started!")
	if loadedConfig.UseDashboard {
		dashboard.Start()
	} else {
		input := bufio.NewScanner(os.Stdin)
		input.Scan()
	}
}
