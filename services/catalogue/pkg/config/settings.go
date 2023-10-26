package config

import (
	"github.com/spf13/viper"
)

func GetSettings(path string) (*Settings, error) {
	viper.SetConfigFile(path)
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	settings := &Settings{}
	err = viper.Unmarshal(settings)
	if err != nil {
		return nil, err
	}

	return settings, nil
}
