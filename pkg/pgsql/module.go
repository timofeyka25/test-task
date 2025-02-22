package pgsql

import "go.uber.org/fx"

var Module = fx.Module("pgsql_connection",
	fx.Provide(
		NewPgsqlConnection,
	),
)
