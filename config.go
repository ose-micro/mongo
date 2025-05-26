package mongodb

import (
	"time"
)

type Config struct {
	Timeout  time.Duration `mapstructure:"timeout"`
	Host     string        `mapstructure:"host"`
	Port     int           `mapstructure:"port"`
	User     string        `mapstructure:"user"`
	Database string        `mapstructure:"database"`
	Password string        `mapstructure:"password"`
}
