package mongo

import (
	"time"
)

type Config struct {
	Timeout  time.Duration
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Database string `mapstructure:"database"`
	Password string `mapstructure:"password"`
}
