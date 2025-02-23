package container

import (
	"go.uber.org/fx"
	"test-task/internal/config"
	pgsqlrepo "test-task/internal/repository/pgsql"
	"test-task/internal/services"
	"test-task/internal/transport/http"
	"test-task/internal/transport/http/handlers"
	"test-task/pkg/pgsql"
)

func Build() *fx.App {
	return fx.New(
		config.Module,
		pgsql.Module,
		http.Module,
		handlers.Module,
		pgsqlrepo.Module,
		services.Module,
	)
}
