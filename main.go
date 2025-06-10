package main

import (
	"api-heartbeat/pkg/config"
	"api-heartbeat/pkg/executor"
	"api-heartbeat/pkg/scheduler"
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type JobConfig struct {
	Name        string
	Enabled     bool
	Frequency   string
	Handler     string
	Description string
}

func loadJobs(jobDirectory string) ([]JobConfig, error) {
	var jobs []JobConfig

	files, err := os.ReadDir(jobDirectory)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		jobConfig, err := loadJobConfig(file.Name())
		if err != nil {
			return nil, err
		}

		jobs = append(jobs, jobConfig)
	}

	return jobs, nil
}

func loadJobConfig(filename string) (JobConfig, error) {
	var jobConfig JobConfig

	// Read the file
	data, err := os.ReadFile(filepath.Join("pkg/jobs", filename))
	if err != nil {
		return JobConfig{}, err
	}

	// Unmarshal the YAML data into the JobConfig struct
	err = yaml.Unmarshal(data, &jobConfig)
	if err != nil {
		return JobConfig{}, err
	}

	return jobConfig, nil
}

func main() {
	cfg := config.InitConfig()
	scheduler, err := scheduler.New(cfg.Timezone)
	if err != nil {
		log.Fatalf("Error creating scheduler: %v", err)
	}

	// Directory for loading jobs
	const jobDirectory = "pkg/jobs"

	// Load and schedule jobs
	jobs, err := loadJobs(jobDirectory)
	if err != nil {
		log.Fatalf("Error loading jobs: %v", err)
	}

	// Schedule each job
	for _, job := range jobs {
		if !job.Enabled {
			continue
		}

		_, err := scheduler.Cron(job.Frequency).Do(func() {
			log.Printf("[Runner] Running job: %s", job.Name)
			err := executor.ExecuteJob(job.Handler, cfg)
			if err != nil {
				log.Printf("[Runner] Error executing job %s: %v", job.Name, err)
			} else {
				log.Printf("[Runner] Job %s executed successfully", job.Name)
			}
		})
		if err != nil {
			log.Printf("[Runner] Error scheduling job %s: %v", job.Name, err)
		}
		log.Printf("[Runner] Job %s scheduled successfully per %s", job.Name, job.Frequency)
	}

	// Start the scheduler
	scheduler.StartBlocking()
}
