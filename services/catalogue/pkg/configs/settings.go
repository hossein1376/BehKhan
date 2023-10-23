package configs

import (
	"github.com/spf13/viper"
)

func GetSettings() (*Settings, error) {
	viper.SetConfigFile("./config/config.yaml")
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
