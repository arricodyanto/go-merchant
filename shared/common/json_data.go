package common

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

var path = "repository/json/"

func ReadJsonFile(fileName string, data interface{}) error {
	file, err := os.ReadFile(filepath.Join(path, fileName))
	if err != nil {
		return fmt.Errorf("failed to read file: %v", err.Error())
	}

	// ubah ke json data
	err = json.Unmarshal(file, &data)
	if err != nil {
		return fmt.Errorf("failed to unmarshal file: %v", err.Error())
	}
	return nil
}

func CreateJsonFile(fileName string) error {
	file, err := os.Create(filepath.Join(path, fileName))
	if err != nil {
		return fmt.Errorf("failed to create customers file: %v", err.Error())
	}
	defer file.Close()

	return nil
}
