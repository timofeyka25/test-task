package pgsql

import "go.uber.org/fx"

var Module = fx.Module("pgsql_repositories",
	fx.Provide(
		NewAuthRepository,
		NewRecordRepository,
	),
)
