package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command
var rootCmd = &cobra.Command{
	Use:   "insertFlag",
	Short: "path and number of goroutine",
	Long:  "A CLI application to perform path of outputFile and number of goroutine.",
	Run: func(cmd *cobra.Command, args []string) {
		log.Debug("Welcome to Mario CLI. Use --help for usage.")

	},
}

// Execute initializes the CLI
func Execute() {

	if err := rootCmd.Execute(); err != nil {
		log.Error(err)
		// os.Exit(1)
	}
}
