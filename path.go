package tiletools

import (
	"log"
	"os"
)

func mkdirIfNotExisting(path string) {
	if pathDontExist(path) {
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			log.Fatalf("Could not create %s: %s", path, err)
		}
	}
}

func pathDontExist(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return true
	}
	return false
}
