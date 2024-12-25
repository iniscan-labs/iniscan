package config

import (
	"log/slog"
	"os"

	"github.com/mss-boot-io/mss-boot/pkg/config"
	"github.com/mss-boot-io/mss-boot/pkg/config/gormdb"
	"github.com/mss-boot-io/mss-boot/pkg/config/source"
)

var Cfg Config

type Config struct {
	Server   config.Listen    `yaml:"server" json:"server"`
	Logger   config.Logger    `yaml:"logger" json:"logger"`
	Database *gormdb.Database `yaml:"database" json:"database"`
}

// Init init config
func (e *Config) Init() {
	opts := []source.Option{
		source.WithDir("config"),
		source.WithProvider(source.Local),
	}
	err := config.Init(e, opts...)
	if err != nil {
		slog.Error("Config init failed", slog.Any("err", err))
		os.Exit(-1)
	}

	e.Logger.Init()
	if e.Database != nil {
		e.Database.Init()
	}
}

// OnChange watch config change
func (e *Config) OnChange() {
	slog.Info("!!! cfg change and reload")
	e.Logger.Init()
	if e.Database != nil {
		e.Database.Init()
	}
}
