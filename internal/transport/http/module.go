package http

import (
	"context"
	"fmt"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

var Module = fx.Module("http",
	fx.Provide(
		fx.Annotate(
			NewServer,
			fx.ParamTags(``, `group:"routes"`),
		),
	),
	fx.Invoke(
		RunServer,
	),
)

func RunServer(lc fx.Lifecycle, server *Server) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				zap.S().Info("Starting HTTP server...")
				if err := server.app.Listen(fmt.Sprintf("%s:%d", server.cfg.Host, server.cfg.Port)); err != nil {
					zap.S().Error("Error starting server", zap.Error(err))
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			zap.S().Info("Shutting down HTTP server...")
			if err := server.app.Shutdown(); err != nil {
				zap.S().Error("Server shutdown error", zap.Error(err))
				return err
			}
			zap.S().Info("Server successfully stopped.")
			return nil
		},
	})
}
