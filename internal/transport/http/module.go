package http

import (
	"context"
	"fmt"
	"go.uber.org/fx"
	"log"
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
				addr := fmt.Sprintf("%s:%d", server.cfg.Host, server.cfg.Port)
				log.Printf("Starting HTTP server at addr %s...", addr)

				if err := server.app.Listen(addr); err != nil {
					log.Println("Error starting server:", err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Println("Shutting down HTTP server...")
			if err := server.app.Shutdown(); err != nil {
				log.Println("Server shutdown error:", err)
				return err
			}
			log.Println("Server successfully stopped.")
			return nil
		},
	})
}
