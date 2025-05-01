package cmd

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/robfig/cron/v3"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type CronJob struct {
	Name     string   `mapstructure:"name"`
	Schedule string   `mapstructure:"schedule"`
	Command  string   `mapstructure:"command"`
	Args     []string `mapstructure:"args"`
	Enabled  bool     `mapstructure:"enabled"`
}

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run configured cron jobs",
	Long: `Run all enabled cron jobs according to their schedules.
The jobs are read from the configuration file.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Initialize cron scheduler
		c := cron.New()

		// Get jobs from configuration
		var jobs []CronJob
		if err := viper.UnmarshalKey("jobs", &jobs); err != nil {
			log.Fatalf("Error reading jobs configuration: %v", err)
		}

		// Add each job to the scheduler
		for _, job := range jobs {
			if !job.Enabled {
				continue
			}

			// Create a closure to capture the job details
			jobFunc := func(j CronJob) func() {
				return func() {
					log.Printf("Running job: %s\n", j.Name)

					// Execute the command
					cmd := exec.Command(j.Command, j.Args...)
					output, err := cmd.CombinedOutput()
					if err != nil {
						log.Printf("Error running job %s: %v\nOutput: %s", j.Name, err, string(output))
						return
					}

					log.Printf("Job %s completed successfully\nOutput: %s\n", j.Name, string(output))
				}
			}(job)

			// Add the job to the scheduler
			_, err := c.AddFunc(job.Schedule, jobFunc)
			if err != nil {
				log.Printf("Error scheduling job %s: %v", job.Name, err)
				continue
			}

			fmt.Printf("Scheduled job: %s with schedule: %s\n", job.Name, job.Schedule)
		}

		// Start the cron scheduler
		c.Start()
		fmt.Println("Cron scheduler started. Press Ctrl+C to stop.")

		// Keep the program running
		select {}
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
