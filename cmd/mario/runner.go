package main

import (
	"bytes"
	"encoding/csv"
	"math/rand"
	"os"
	"strconv"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
)

func randomString(length int, fixStr bool) string {
	if fixStr {
		counter = counter + 1
		strNumber := strconv.FormatUint(counter, 10)
		mouStr := "A" + strNumber
		return mouStr
	}
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

	// Keep adding to the map until it reaches up 1MB in size
	for {
		// 1 second delay delay
		time.Sleep(1000 * time.Millisecond)
		mu.Lock()
		if getMapSize(data) >= targetSize {
			mu.Unlock()
			break
		}

		name := randomString(lenRndStr, fixString)
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

func statFile(filePath string) error {

	// Check if the file exists
	_, err := os.Stat(filePath)
	if err == nil {
		// File exists, remove it
		err = os.Remove(filePath)
		if err != nil {
			log.Error("Error removing file:", err)
			return err
		}
		log.Debug("File " + filePath + " removed successfully.")
	} else if os.IsNotExist(err) {
		// File does not exist, handle accordingly
		log.Debug("File does not exist. ok continue")
		return nil
	} else {
		// Other error occurred
		log.Error("Error checking file:", err)
		return err
	}
	return nil
}

func runner(numGoroutines int, pathCsv string, lenStringRnd int, NumberRnd int) error {
	// Map to store the generated names and numbers
	data := make(map[string]int)
	var mu sync.Mutex
	var wg sync.WaitGroup

	// Start goroutines
	log.Info("number of Goroutines:", numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go populateMap(data, &mu, &wg, lenStringRnd, NumberRnd)
	}

	// Wait for all goroutines to finish
	wg.Wait()

	// Create a CSV file
	file, err := os.Create(pathCsv)
	if err != nil {
		log.Error("Error creating CSV file:", err)
		return err
	}
	defer file.Close()

	// Create a CSV writer
	writer := csv.NewWriter(file)

	// Write the map data to the CSV
	for key, value := range data {
		err := writer.Write([]string{key, strconv.Itoa(value)})
		if err != nil {
			log.Error("Error writing to CSV file:", err)
			return err
		}
	}

	// Flush the writer to ensure data is written
	writer.Flush()

	// Check if there was any error during the writing
	if err := writer.Error(); err != nil {
		log.Error("Error during writing CSV:", err)
	}

	log.Info("Data written to successfully!")
	return nil
}
