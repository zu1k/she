package executor

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

// forward compatibility before 1.0
func readRawConfig(path string) ([]byte, error) {
	data, err := ioutil.ReadFile(path)
	if err == nil && len(data) != 0 {
		return data, nil
	}

	if filepath.Ext(path) != ".yaml" {
		return nil, err
	}

	path = path[:len(path)-5] + ".yml"
	if _, fallbackErr := os.Stat(path); fallbackErr == nil {
		return ioutil.ReadFile(path)
	}

	return data, err
}

func readConfig(path string) ([]byte, error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil, err
	}
	data, err := readRawConfig(path)
	if err != nil {
		return nil, err
	}

	if len(data) == 0 {
		return nil, fmt.Errorf("Configuration file %s is empty", path)
	}

	return data, err
}
