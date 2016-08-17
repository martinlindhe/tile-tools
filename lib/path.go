package tiletools

import "os"

// PathDontExist returns true if path does not exist
func PathDontExist(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return true
	}
	return false
}
