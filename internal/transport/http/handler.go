package http

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

type Handler interface {
	Register(router fiber.Router)
}

func AsHandler(f any) any {
	return fx.Annotate(
		f,
		fx.As(new(Handler)),
		fx.ResultTags(`group:"routes"`),
	)
}
