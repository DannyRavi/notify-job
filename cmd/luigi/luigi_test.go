package main

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

var status bool

func task1(ctx context.Context) {
	// Simulate a long-running task
	time.Sleep(3 * time.Second)
	fileName := "/tmp/check_test.txt"
	err := statFile(fileName)
	if err != nil { //fixme: change to error type
		log.Error("Error:", err)
		return
	}
	data := []byte("This is the content of the file.\n")

	err = os.WriteFile(fileName, data, 0644)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}

	// fmt.Println("File created successfully:", fileName)
	// fmt.Println("Task 1 completed")
}

func task2() {
	// Simulate a short-running task
	// time.Sleep(3 * time.Second)
	dirName := "/tmp/"
	err := executor(dirName)
	if err != nil { //fixme: change to error type
		status = false
		log.Error("Error:", err)
		return
	}
	status = true
	// fmt.Println("Task 2 completed")

}

func TestCheckNotify(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	go task1(ctx)
	task2()

	<-ctx.Done()
	// fmt.Println(status)
	assert.Equal(t, true, status)
	// assert.Equal(t, nil, result)

}

func TestHash256(t *testing.T) {

	fileName := "/tmp/check_test.txt"
	data := []byte("This is the content of the file.\n")

	err := os.WriteFile(fileName, data, 0644)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	file, _ := os.Open(fileName)
	hash, err := calculateHash(file)
	if err != nil {
		log.Error("Error calculating hash:", err)
		return
	}
	hashRec := fmt.Sprintf("%x", hash)
	hashSave := "72b19a6579b1a278d5730f857c2758d0bf5b948718f11d5835dccbe5b10e6f69"
	assert.Equal(t, hashRec, hashSave)
	// assert.Equal(t, nil, result)

}

// func Add(x, y int) int {
// 	return x + y
// }

// func TestAdd2(t *testing.T) {
// 	assert.Equal(t, 5, Add(2, 3))
// }
