package main

import (
	"bufio"
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
)

func showLogo() {
	fileName := "../../assets/logo_Luigi.txt"

	// Open the file
	file, err := os.Open(fileName)
	if err != nil {
		log.Error("Error: Unable to open file: ", fileName, err)
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
		log.Error("Error: Unable to read file: ", fileName, err)
	}
}
