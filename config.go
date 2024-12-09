package raas

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	Server   ServerConfig
	Market   MarketConfig
	Database DatabaseConfig
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

type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Name     string
	SSLMode  string
	Net      string
	Timeout  time.Duration
}

type EnvName string
type AppName string

func Load() (*Config, error) {
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	err = godotenv.Load(filepath.Join(pwd, "../.env"))
	if err != nil {
		fmt.Errorf("Error loading .env file: %w", err)
	}
	viper.SetConfigFile("../config/config.yaml")

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := viper.ReadInConfig(); err != nil {
		log.Println("Warning: Config file not found, using environment variables only.")
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
