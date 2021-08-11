package rdb

import (
	"time"
)

type Config struct {
	DSN             string
	ConnMaxLifetime time.Duration
	MaxIdleConns    int
	SetMaxOpenConns int
}
