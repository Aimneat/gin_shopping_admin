package config

import "time"

type Server struct {
	RunMode      string
	Host         string
	Port         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}
