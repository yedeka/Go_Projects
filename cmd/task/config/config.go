package config

import (
	"fmt"
	"path/filepath"

	homeDir "github.com/mitchellh/go-homedir"
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
	// Use User's home directory to create the databse
	homedirectory, err := homeDir.Dir()
	if nil != err {
		return nil, fmt.Errorf(" error while locating home directory of user")
	}
	dbPath := filepath.Join(homedirectory, config.Database.DbName)
	config.Database.DbName = dbPath
	return config, nil
}
