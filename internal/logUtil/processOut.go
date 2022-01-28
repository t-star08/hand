package logUtil

import (
	"fmt"
	"os"
)

func CheckLogDirectory(logDirectoryPath string) error {
	if file, err := os.Stat(logDirectoryPath); os.IsNotExist(err) || !file.IsDir() {
		return fmt.Errorf("log directory not exist")
	}
	return nil
}

func CheckLogFile(logFilePath string) error {
	if _, err := os.Stat(logFilePath); os.IsNotExist(err) {
		return fmt.Errorf("log file not exist")
	}
	return nil
}

func CreateLogFile(logFilePath string) error {
	if _, err := os.Create(logFilePath); err != nil {
		return err
	}
	return nil
}
