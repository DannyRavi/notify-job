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

	fmt.Println("Running as root...")
	start := time.Now()
	Execute()
	duration := time.Since(start)

	// Print the execution time
	fmt.Printf("Execution time: %v\n", duration)
}
