package http

import "time"

type Config struct {
	Host         string
	Port         uint16
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}
