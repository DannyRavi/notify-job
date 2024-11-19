package main

import (
	"time"

	log "github.com/sirupsen/logrus"
)

var (
	verbose           bool
	outputPath        string
	concurrent        int
	minLimitGoroutine int = 1     // Minimum allowed value
	maxLimitGoroutine int = 10000 // Maximum allowed value

	setLenRandomString int = 70     // Maximum len name
	setRandomNumber    int = 1000    // Maximum random number
	targetSize         int = 1 << 20 // 1MB
	// targetSize = 1 << 20 // 1MB
	// targetSize = 1 << 17 // 1MB
)

// const targetSize = 1 << 4 // under 1MB
func init() {
	showLogo()
	log.SetFormatter(&log.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})
	// Log as JSON instead of the default ASCII formatter.
	// log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.InfoLevel)

	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringVarP(&outputPath, "path", "d", "/tmp/output.csv", "Path to save the output file")
	addCmd.Flags().IntVarP(&concurrent, "concurrent", "r", 50, "number of concurrent")
	addCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "Enable verbose mode")
	addCmd.MarkFlagRequired("path")
	addCmd.MarkFlagRequired("concurrent")
	// addCmd.MarkFlagRequired("verbose")

}

func main() {
	start := time.Now()
	Execute()
	duration := time.Since(start)

	// Print the execution time
	log.Info("Execution time: ", duration)

}
