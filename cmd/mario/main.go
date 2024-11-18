package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	Execute()
	duration := time.Since(start)

	// Print the execution time
	fmt.Printf("Execution time: %v\n", duration)
}
