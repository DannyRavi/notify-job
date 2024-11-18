package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command
var rootCmd = &cobra.Command{
	Use:   "insertFlag",
	Short: "path and number of goroutine",
	Long:  "A CLI application to perform path of outputFile and number of goroutine.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to Mario CLI. Use --help for usage.")

	},
}

// Execute initializes the CLI
func Execute() {
	showLogo()
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
