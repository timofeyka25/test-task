package main

import (
	"go.uber.org/zap"
	"test-task/internal/container"
	"time"
)

func main() {
	now := time.Now()

	app := container.Build()

	zap.S().Info("Starting application...")

	zap.S().Infof("Up and running (%s)", time.Since(now))

	app.Run()

	zap.S().Info("Service stopped.")
}
