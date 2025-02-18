package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command
var rootCmd = &cobra.Command{
	Use:   "insertFlag",
	Short: "path and number of goroutine",
	Long:  "A CLI application to monitor path of outputFile and PID of that.",
	Run: func(cmd *cobra.Command, args []string) {
		log.Debug("Welcome to Luigi CLI. Use --help for usage.")

	},
}

// Execute initializes the CLI
func Execute() {

	if err := rootCmd.Execute(); err != nil {
		log.Panic(err)
	}
}
