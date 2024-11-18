package main

import (
	"bufio"
	"fmt"
	"os"
)

func showLogo() {
	fileName := "../../assets/logo_Mario.txt"

	// Open the file
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("Error: Unable to open file %s: %v\n", fileName, err)
		return
	}
	defer file.Close()

	// Read the file line by line
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	// Check for errors in scanning
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error: Unable to read file %s: %v\n", fileName, err)
	}
}
