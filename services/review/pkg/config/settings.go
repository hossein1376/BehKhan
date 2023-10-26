package config

import (
	"encoding/json"
	"fmt"
	"os"
)

func GetConfigs(path string) (*Settings, error) {
	if path == "" {
		return nil, fmt.Errorf("empty path for config file")
	}

	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	settings := new(Settings)
	err = json.Unmarshal(file, settings)
	if err != nil {
		return nil, err
	}

	return settings, nil
}
