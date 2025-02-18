package main

import (
	"crypto/sha256"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/opcoder0/fanotify"
	log "github.com/sirupsen/logrus"
)

func calculateHash(reader io.Reader) ([]byte, error) {
	h := sha256.New()
	if _, err := io.Copy(h, reader); err != nil {
		return nil, err
	}
	return h.Sum(nil), nil
}

func openFiles(filaPath string) {
	// Using a file as input
	file, err := os.Open(filaPath)
	if err != nil {
		log.Error("Error opening file:", err)
		return
	}
	defer file.Close()

	hash, err := calculateHash(file)
	if err != nil {
		log.Error("Error calculating hash:", err)
		return
	}

	log.Info("SHA256 hash: ", fmt.Sprintf("%x", hash))
	removeFile(filaPath)
	return
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

func removeFile(filePathRemove string) {

	err := os.Remove(filePathRemove)
	if err != nil {
		log.Error("Error removing file:", err)
		return
	}

	log.Debug("File removed successfully.")
}

func executor(watchDir string) error {
	// Initialize fanotify
	l, err := fanotify.NewListener("/", false, fanotify.PermissionNone)
	if err != nil {
		log.Error("Event use ---> sudo <---: ", err)
		return err
	}

	// t.Logf("Watch Directory: %s", watchDir)
	eventType := fanotify.FileCreated
	l.AddWatch(watchDir, eventType)
	go l.Start()
	defer l.Stop()

	select {
	case <-time.After(3600 * time.Second):
		log.Error("Timeout Error: FileOpenedForExec event not received")
	case event := <-l.Events:

		log.Info("File path: ", fmt.Sprintf("%s/%s", event.Path, event.FileName))
		log.Info("PID: ", event.Pid)
		log.Debug("file created? ", event.EventTypes.Has(fanotify.FileCreated))
		fullPath := event.Path + "/" + event.FileName
		openFiles(fullPath)

	}
	return nil
}
