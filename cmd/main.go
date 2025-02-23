package main

import (
	"log"
	"test-task/internal/container"
	"time"
)

func main() {
	now := time.Now()

	app := container.Build()

	log.Println("Starting application...")

	log.Printf("Up and running (%s)", time.Since(now))

	app.Run()

	log.Println("Service stopped.")
}
