package config

import (
	"time"
)

type JWTConfig struct {
	Key    string
	Expire time.Duration
}
