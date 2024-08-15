package util

import "os"

func ReadFile(path string) ([]byte, error) {
	fileContents, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return fileContents, nil
}
