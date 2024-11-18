package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "cli",
	Short: "write random number to file",
	Long:  "Perform addition write random map to file",

	Run: func(cmd *cobra.Command, args []string) {
		if concurrent < minLimitGoroutine || concurrent > maxLimitGoroutine {
			panic("please insert concurrent (r) number between 1 to 100")
		}
		err := runner(concurrent, outputPath, setLenRandomString, setRandomNumber)
		fmt.Print(concurrent)
		fmt.Print(outputPath)
		if err != nil { //fixme: change to error type
			fmt.Println("Error:", err)
			return
		}
		// fmt.Printf("Result of addition: %.2f\n", result)
	},
}

var (
	outputPath        string
	concurrent        int
	minLimitGoroutine int = 1    // Minimum allowed value
	maxLimitGoroutine int = 1000 // Maximum allowed value

	setLenRandomString int = 10    // Maximum len name
	setRandomNumber    int = 10000 // Maximum random number
)

const targetSize = 1 << 18 // 1MB
// const targetSize = 1 << 20 // 1MB
func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringVarP(&outputPath, "path", "d", "output.csv", "Path to save the output file")
	addCmd.Flags().IntVarP(&concurrent, "concurrent", "r", 50, "number of concurrent")
	// addCmd.Flags().Float64VarP(&num2, "num2", "b", 0, "Second number (required)")
	addCmd.MarkFlagRequired("path")
	addCmd.MarkFlagRequired("concurrent")

}
