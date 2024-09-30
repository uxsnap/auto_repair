package config

import "os"

type ConfigHTTP struct {
	Host string
	Port string
}

func NewConfigHttp() *ConfigHTTP {
	return &ConfigHTTP{
		Host: os.Getenv("SERVICE_HOST"),
		Port: os.Getenv("SERVICE_PORT"),
	}
}
