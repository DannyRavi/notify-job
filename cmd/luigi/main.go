package main

import (
	"os"
	"time"

	log "github.com/sirupsen/logrus"
)

var (
	verbose       bool
	directoryPath string
)

func init() {

	log.SetFormatter(&log.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})
	log.SetLevel(log.InfoLevel)

	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringVarP(&directoryPath, "path", "d", "/tmp", "Path to watch the directory (please same address of Mario App)")
	addCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "Enable verbose mode")
	addCmd.MarkFlagRequired("path")
	// addCmd.MarkFlagRequired("verbose")

}

func main() {
	showLogo()
	if os.Geteuid() != 0 {
		log.Panic("Error: This program must be run as root.")
	}

	pwd, err := os.Getwd()
	if err != nil {
		log.Error("Error getting working directory:", err)
		return
	}

	log.Debug("Current working directory:", pwd)

	log.Info("Running as root...")

	start := time.Now()
	Execute()
	duration := time.Since(start)

	// Print the execution time
	log.Info("Execution time:", duration)
}
