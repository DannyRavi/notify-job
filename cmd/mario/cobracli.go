package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "cli",
	Short: "write random number to file",
	Long:  "Perform addition write random map to file",

	Run: func(cmd *cobra.Command, args []string) {
		log.Info("number of goroutines: ", concurrent)
		log.Info("output file path: ", outputPath)
		if verbose {
			log.SetLevel(log.DebugLevel)
			log.Debug("verbose mode: ", verbose)
		}
		if concurrent < minLimitGoroutine || concurrent > maxLimitGoroutine {
			log.Panic("please insert concurrent (r) number between 1 to 10000")
		}
		err := statFile(outputPath)
		if err != nil { //fixme: change to error type
			log.Error("Error:", err)
			return
		}
		err = runner(concurrent, outputPath, setLenRandomString, setRandomNumber)
		if err != nil { //fixme: change to error type
			log.Error("Error:", err)
			return
		}
	},
}
