package raas

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Server ServerConfig
	Market MarketConfig
}

type ServerConfig struct {
	Port           int
	DebugPort      int
	Production     bool
	HandlerTimeout time.Duration `validate:"min=1"`
}

type MarketConfig struct {
	PSX PsxConfig
}

type PsxConfig struct {
	BaseURL       string
	ScraperURL    ScraperURL
	TimeseriesURL TimeseriesURL
}
type ScraperURL struct {
	Indices string
	Company string
	Reports string
	Symbols string
}

type TimeseriesURL struct {
	EOD      string
	Intraday string
}

type EnvName string
type AppName string

// Load reads and parses the configuration file, returning a Config instance.
func Load() (*Config, error) {
	viper.SetConfigName("config") // Name of the config file (without extension)
	viper.SetConfigType("yaml")   // File format (yaml, json, etc.)
	// For Debug mode "../config/", otherwise "config/"
	viper.AddConfigPath("../config/") // Look for the config file in the `config` directory
	viper.AutomaticEnv()              // Override config with environment variables if they exist

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
