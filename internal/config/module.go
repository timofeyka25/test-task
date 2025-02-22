package config

import (
	"go.uber.org/fx"
	"test-task/internal/transport/http"
	"test-task/pkg/pgsql"
)

var Module = fx.Module("config",
	fx.Provide(
		New,
		func(cfg *Config) *pgsql.Config {
			return cfg.PgSQLConfig
		},

		func(cfg *Config) *http.Config {
			return cfg.HTTPConfig
		},
	),
)
