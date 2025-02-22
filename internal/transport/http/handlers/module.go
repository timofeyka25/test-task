package handlers

import (
	"go.uber.org/fx"
	"test-task/internal/transport/http"
)

var Module = fx.Module("handlers",
	fx.Provide(
		http.AsHandler(NewMetaHandler),
	),
)
