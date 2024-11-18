package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"sync"
	"time"
)

func randomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	rand.NewSource(time.Now().UnixNano())
	var result []byte
	for i := 0; i < length; i++ {
		result = append(result, charset[rand.Intn(len(charset))])
	}
	return string(result)
}

// Goroutine function to populate the map
func populateMap(data map[string]int, mu *sync.Mutex, wg *sync.WaitGroup, lenRndStr int, RndNum int) {
	defer wg.Done()

	// Keep adding to the map until it reaches 1MB in size
	for {
		// add delay
		mu.Lock()
		if getMapSize(data) >= targetSize {
			mu.Unlock()
			break
		}

		name := randomString(lenRndStr)
		number := rand.Intn(RndNum)
		data[name] = number
		mu.Unlock()
	}
}

// Calculate the size of the map in bytes
func getMapSize(data map[string]int) int {
	var buf bytes.Buffer
	for key, value := range data {
		buf.WriteString(key)
		buf.WriteString(strconv.Itoa(value))
	}
	return buf.Len()
}

func runner(numGoroutines int, pathCsv string, lenStringRnd int, NumberRnd int) error {
	// Map to store the generated names and numbers
	data := make(map[string]int)
	var mu sync.Mutex
	var wg sync.WaitGroup

	// Start goroutines
	fmt.Println("number of runner:", numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go populateMap(data, &mu, &wg, lenStringRnd, NumberRnd)
	}

	// Wait for all goroutines to finish
	wg.Wait()

	// Create a CSV file
	file, err := os.Create(pathCsv)
	if err != nil {
		fmt.Println("Error creating CSV file:", err)
		return err
	}
	defer file.Close()

	// Create a CSV writer
	writer := csv.NewWriter(file)

	// Write the map data to the CSV
	for key, value := range data {
		err := writer.Write([]string{key, strconv.Itoa(value)})
		if err != nil {
			fmt.Println("Error writing to CSV file:", err)
			return err
		}
	}

	// Flush the writer to ensure data is written
	writer.Flush()

	// Check if there was any error during the writing
	if err := writer.Error(); err != nil {
		fmt.Println("Error during writing CSV:", err)
	}

	fmt.Println("Data written to output.csv successfully!")
	return nil
}
