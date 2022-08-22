package main

import (
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/help-14/ocean-backup/config"
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
		log.Fatalln("Reading config.yaml failed!", err)
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
			log.Fatalln("Unsupported service name: "+serviceConfig.Service, err)
		}

		err = newService.Setup(serviceConfig)
		if err != nil {
			log.Fatalln("Setup service '"+serviceConfig.Name+"' failed!", err)
		}
		err = newService.Connect()
		if err != nil {
			log.Fatalln("Connect to service '"+serviceConfig.Name+"' failed!", err)
		}
		log.Println("Service '" + serviceConfig.Name + "' is connected!")
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
			runner.Run()
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

	// Hello world, the web server
	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello, world!\n")
	}

	http.HandleFunc("/hello", helloHandler)
	log.Println("Ocean backup is started!")
	log.Println("Web dashboard is running at http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
