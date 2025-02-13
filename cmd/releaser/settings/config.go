package settings

import (
	"os"
	"time"
)

const (
	DefaultPoolTimeout = 30 * time.Second
)

type Config struct {
	Telegram    *Telegram
	PoolTimeout time.Duration
}

type Telegram struct {
	Token string
	// Admin - admin telegram_id
	Admin string
}

func NewConfig() *Config {
	c := &Config{
		PoolTimeout: DefaultPoolTimeout, Telegram: &Telegram{},
	}

	return c
}

func (c *Config) WithEnv() error {
	if botToken := os.Getenv("TELEGRAM_BOT_TOKEN"); botToken != "" {
		c.Telegram.Token = botToken
	}
	if admin := os.Getenv("TELEGRAM_ADMIN"); admin != "" {
		c.Telegram.Admin = admin
	}
	if poolTimeout := os.Getenv("POOL_TIMEOUT"); poolTimeout != "" {
		c.PoolTimeout, _ = time.ParseDuration(poolTimeout)
	}

	return nil
}
