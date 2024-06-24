package config

import (
	"errors"
	"fmt"
	"path"
	"time"

	"github.com/spf13/viper"
)

var (
	ErrInvalidConfigPath = errors.New("invalid config path")
)

type Config interface {
	String(key string) string
	StringSlice(key string) []string
	Int(key string) int
	Float(key string) float64
	Bool(key string) bool
	Time(key string) time.Time
	Duration(key string) time.Duration
}

type config struct {
	cfg *viper.Viper
}

func New(envPath string) (Config, error) {
	cfg := viper.New()
	cfg.SetConfigFile(path.Join(envPath))

	if err := cfg.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("%w: %s", ErrInvalidConfigPath, err)
	}

	cfg.WatchConfig()
	cfg.AutomaticEnv()

	return config{
		cfg: cfg,
	}, nil
}

func (c config) String(key string) string {
	return c.cfg.GetString(key)
}

func (c config) StringSlice(key string) []string {
	return c.cfg.GetStringSlice(key)
}

func (c config) Int(key string) int {
	return c.cfg.GetInt(key)
}

func (c config) Float(key string) float64 {
	return c.cfg.GetFloat64(key)
}

func (c config) Bool(key string) bool {
	return c.cfg.GetBool(key)
}

func (c config) Time(key string) time.Time {
	return c.cfg.GetTime(key)
}

func (c config) Duration(key string) time.Duration {
	return c.cfg.GetDuration(key)
}
