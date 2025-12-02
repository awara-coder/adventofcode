package utils

import (
	"bufio"
	"os"
	"path/filepath"
)

func ReadFileContents(assetFilePath string) ([]string, error) {
	currentWorkingDirectory, err := os.Getwd()
	if err != nil {
		GetLogger().Fatal("error while fetching current working directory")
		return nil, err
	}
	file, err := os.Open(filepath.Join(currentWorkingDirectory, "../assets", assetFilePath))
	if err != nil {
		GetLogger().Printf("Unable to read file: path: %s, error: %s\n", assetFilePath, err.Error())
		return nil, err
	}
	defer file.Close()

	lines := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, nil

}
