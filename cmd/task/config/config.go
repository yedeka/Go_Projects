package config

import (
	"github.com/spf13/viper"
)

type Configurations struct {
	Server   ServerConfig
	Database DatabaseConfiguration
}

type ServerConfig struct {
	Port int
}

type DatabaseConfiguration struct {
	DbName       string
	DbTimeout    int
	DbBucketName string
}

func setConfigDefaults() {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
}

func LoadConfigurations() (*Configurations, error) {
	setConfigDefaults()
	err := viper.ReadInConfig()

	if nil != err {
		return nil, err
	}
	var config *Configurations
	err = viper.Unmarshal(&config)

	if nil != err {
		return nil, err
	}
	return config, nil
}
