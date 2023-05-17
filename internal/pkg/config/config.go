package config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type DatabaseConfiguration struct {
	DatabaseType string `envconfig:"DATABASE_TYPE" default:"mysql"`
	Username     string `envconfig:"USERNAME" default:"root"`
	Password     string `envconfig:"PASSWORD" default:"root"`
	Port         int    `envconfig:"PORT" default:"3306"`
	DbName       string `envconfig:"DATABASE_NAME" default:"testing"`
	Url          string `envconfig:"URL" default:"127.0.0.1"`
}

func GetDatabaseConfig() (*DatabaseConfiguration, error) {
	var databaseConf DatabaseConfiguration
	if err := envconfig.Process("", &databaseConf); err != nil {
		return nil, fmt.Errorf("configuration failed %v", err)
	}

	return &databaseConf, nil
}
