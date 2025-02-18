package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateAndeCheck(t *testing.T) {
	// targetSize = 1 << 10 // 1MB
	targetSize = 1 << 10
	fixString = true
	concurrent := 100
	outputPath := "/tmp/outs.csv"
	setLenRandomString := 10
	setRandomNumber := 10
	err := runner(concurrent, outputPath, setLenRandomString, setRandomNumber)
	if err != nil {
		// log.Error("Error:", err)
		fmt.Println(err)
		return
	}

	// Open the file
	file, err := os.Open(outputPath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println("Error getting file info:", err)
		return
	}

	// Get file size in bytes
	fileSizeInBytes := fileInfo.Size()

	fmt.Println("File size:", fileSizeInBytes, "bytes")
	const recSize int64 = 1481
	assert.Equal(t, recSize, fileSizeInBytes)
	// sum := Add(2, 3)
	// if sum != 5 {
	// 	t.Errorf("Add(2, 3) = %d; want 5", sum)
	// }
}

func Add(x, y int) int {
	return x + y
}

func TestAdd2(t *testing.T) {
	assert.Equal(t, 5, Add(2, 3))
}
