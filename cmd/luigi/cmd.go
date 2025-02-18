package main

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "cli",
	Short: "Monitor direcory",
	Long:  "Monitor direcory then find create file and remove it",

	Run: func(cmd *cobra.Command, args []string) {
		if verbose {
			log.SetLevel(log.DebugLevel)
			log.Debug("verbose mode: ", verbose)
		}
		log.Debug("luigi input dir: ", directoryPath)
		err := executor(directoryPath)
		log.Debug(directoryPath)

		if err != nil { //fixme: change to error type
			fmt.Println("Error:", err)
			return
		}
		// fmt.Printf("Result of addition: %.2f\n", result)
	},
}
