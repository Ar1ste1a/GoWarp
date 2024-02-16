package util

import "os"

func FileExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	return false
}

func HTBExists(fp string) bool {
	return FileExists(fp)
}

func ReadFile(path string) ([]byte, error) {
	return os.ReadFile(path)
}

func WriteFile(path string, data []byte) error {
	// Write the bytes to file
	err := os.WriteFile(path, data, 0644)
	return err
}
