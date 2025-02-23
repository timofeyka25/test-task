package config

import (
	"fmt"
	"github.com/spf13/viper"
	"sync"
	"test-task/internal/services"
	"test-task/internal/transport/http"
	"test-task/pkg/pgsql"
)

var (
	err    error
	config *Config
	once   sync.Once
)

type Config struct {
	PgSQLConfig  *pgsql.Config
	HTTPConfig   *http.Config
	AccessConfig *services.AccessConfig
}

func New() (*Config, error) {
	once.Do(func() {
		config = &Config{}

		viper.AddConfigPath(".")
		viper.SetConfigName("config")
		if err = viper.ReadInConfig(); err != nil {
			return
		}

		databaseConfig := viper.Sub("db")
		httpConfig := viper.Sub("http")
		accessConfig := viper.Sub("access")

		if err = parseSubConfig(databaseConfig, &config.PgSQLConfig); err != nil {
			return
		}
		if err = parseSubConfig(httpConfig, &config.HTTPConfig); err != nil {
			return
		}
		if err = parseSubConfig(accessConfig, &config.AccessConfig); err != nil {
			return
		}
	})

	return config, err
}

func parseSubConfig[T any](subConfig *viper.Viper, parseTo *T) error {
	if subConfig == nil {
		return fmt.Errorf("can not read %T config: subconfig is nil", parseTo)
	}

	if err := subConfig.Unmarshal(parseTo); err != nil {
		return err
	}

	return nil
}
