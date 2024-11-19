package main

import (
	"fmt"
	"os"
	"time"
)

func main() {

	if os.Geteuid() != 0 {
		fmt.Println("Error: This program must be run as root.")
		os.Exit(1)
	}

	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting working directory:", err)
		return
	}

	fmt.Println("Current working directory:", pwd)

	fmt.Println("Running as root...")
	start := time.Now()
	Execute()
	duration := time.Since(start)

	// Print the execution time
	fmt.Printf("Execution time: %v\n", duration)
}
